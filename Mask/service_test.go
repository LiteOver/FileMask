package Mask_test

import (
	"FileMask/Mask"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func TestProd_Produce(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockProducer := new(MockProducer)

		mockProducer.On("Produce").Return([]string{"ahttp://text .... http://text"}, nil)

		result, err := mockProducer.Produce()

		assert.NoError(t, err)
		assert.Equal(t, []string{"ahttp://text .... http://text"}, result)

		mockProducer.AssertExpectations(t)
	})
}

type MockPresentor struct {
	mock.Mock
}

func (m *MockPresentor) Present(s []string) error {
	args := m.Called(s)
	return args.Error(0)
}

func TestPresent_Present(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockPresentor := new(MockPresentor)
		mockPresentor.On("Present", []string{"ahttp://text .... http://text"}).Return(nil)
		err := mockPresentor.Present([]string{"ahttp://text .... http://text"})
		assert.NoError(t, err)
		mockPresentor.AssertExpectations(t)

	})
}

func TestDataMask(t *testing.T) {
	data := []struct {
		name     string
		text     []string
		expected []string
	}{
		{"first", []string{"http://texttext text"}, []string{"http://******** text"}},
		{"second", []string{"hi http://http:// bye"}, []string{"hi http://******* bye"}},
		{"third", []string{"hi http://te!*xt"}, []string{"hi http://******"}},
	}
	s := Mask.Service{Mask.Prod{}, Mask.Present{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := s.DataMask(d.text)
			fmt.Println(result)
			assert.Equal(t, d.expected, result)
		})
	}

}

func TestService_Run(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mProducer := Mask.NewProducer("/Users/bocman/GolandProjects/FileMask/file.txt")
		mPresenter := Mask.NewPresenter("/Users/bocman/GolandProjects/FileMask/file2.txt")
		workNewService := Mask.NewService(mProducer, mPresenter)
		err := workNewService.Run()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Error", func(t *testing.T) {
		mProducer := Mask.NewProducer("s")
		mPresenter := Mask.NewPresenter("")
		workNewService := Mask.NewService(mProducer, mPresenter)
		workNewService.Run()
		_, err := workNewService.Prod.Produce()
		expected := "open s: no such file or directory"
		if err.Error() != expected {
			t.Errorf("Expected: %s, Got: %s", expected, err.Error())
		}

	})
}
