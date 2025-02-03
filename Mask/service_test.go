package Mask_test

import (
	"FileMask/Mask"
	"github.com/google/go-cmp/cmp"
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

func TestNewService(t *testing.T) {
	expected := &Mask.Service{&Mask.Prod{Adress: "/Users/bocman/GolandProjects/FileMask/file.txt"}, &Mask.Present{"/Users/bocman/GolandProjects/FileMask/find.txt"}}
	result := Mask.NewService(Mask.NewProducer("/Users/bocman/GolandProjects/FileMask/file.txt"), Mask.NewPresenter("/Users/bocman/GolandProjects/FileMask/find.txt"))
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestDataMask(t *testing.T) {
	mProducer := Mask.NewProducer("/Users/bocman/GolandProjects/FileMask/file.txt")
	s := Mask.Service{mProducer, Mask.Present{}}
	expected := []string{"ahttp://**** .... http://****"}
	data, _ := s.Prod.Produce()
	result := s.DataMask(data)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
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
