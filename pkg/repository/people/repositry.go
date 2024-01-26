package people

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	res "people/pkg/repository/people/reponse"
	"sync"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) GetByName(name string) (*res.People, error) {
	var wg sync.WaitGroup

	// Создаем каналы для передачи результатов и ошибок от горутин
	agifyCh := make(chan []byte, 1)
	genderizeCh := make(chan []byte, 1)
	nationalizeCh := make(chan []byte, 1)
	errCh := make(chan error, 3)

	// Инкрементируем счетчик горутин
	wg.Add(3)

	// Запускаем горутины для каждого запроса
	go func() {
		defer wg.Done()
		agifyResult, err := r.getAgifyData(name)
		if err != nil {
			errCh <- err
			return
		}
		agifyCh <- agifyResult
	}()

	go func() {
		defer wg.Done()
		genderizeResult, err := r.getGenderizeData(name)
		if err != nil {
			errCh <- err
			return
		}
		genderizeCh <- genderizeResult
	}()

	go func() {
		defer wg.Done()
		nationalizeResult, err := r.getNationalizeData(name)
		if err != nil {
			errCh <- err
			return
		}
		nationalizeCh <- nationalizeResult
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()

	close(agifyCh)
	close(genderizeCh)
	close(nationalizeCh)
	close(errCh)

	// Считываем результаты из каналов
	agifyResult := <-agifyCh
	genderizeResult := <-genderizeCh
	nationalizeResult := <-nationalizeCh

	// Обрабатываем ошибки
	if len(errCh) > 0 {
		err := <-errCh
		return nil, err
	}

	// Объединяем результаты в один объект
	responseModel := &res.People{}

	if err := json.Unmarshal(agifyResult, responseModel); err != nil {
		fmt.Println("Ошибка при десериализации JSON (Agify):", err)
		return nil, err
	}

	if err := json.Unmarshal(genderizeResult, responseModel); err != nil {
		fmt.Println("Ошибка при десериализации JSON (Genderize):", err)
		return nil, err
	}

	if err := json.Unmarshal(nationalizeResult, responseModel); err != nil {
		fmt.Println("Ошибка при десериализации JSON (Nationalize):", err)
		return nil, err
	}

	return responseModel, nil
}

func (r Repository) getAgifyData(name string) ([]byte, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	return r.makeRequest(url)
}

func (r Repository) getGenderizeData(name string) ([]byte, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	return r.makeRequest(url)
}

func (r Repository) getNationalizeData(name string) ([]byte, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	return r.makeRequest(url)
}

func (r Repository) makeRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке GET-запроса:", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Неправильный статус ответа:", response.Status)
		return nil, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
