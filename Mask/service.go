package Mask

type Service struct {
	Prod Producer
	Pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{Prod: prod, Pres: pres}
}

func (s *Service) DataMask(data []string) []string {
	var prefix string = "http://"
	for i, v := range data {
		b := []byte(v)
		for j := range b {
			if j < len(b)-len(prefix) && string(b[j:j+len(prefix)]) == prefix {
				for j < len(b)-len(prefix) && (string(v[j+len(prefix)]) != " ") {
					mask := "*"
					b[j+len(prefix)] = mask[0]
					j++
				}
			}
		}
		data[i] = string(b)
	}
	return data
}

func (s *Service) Run() error {
	data, err := s.Prod.Produce()
	if err != nil {
		return err
	}

	maskedData := s.DataMask(data)

	err = s.Pres.Present(maskedData)
	if err != nil {
		return err
	}

	return nil
}
