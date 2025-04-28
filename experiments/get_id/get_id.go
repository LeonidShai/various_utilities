package getid

import "fmt"

type GetIdInterface interface {
	GetId() string
}

type Root struct {
	Id             *string
	ParentInstance GetIdInterface
	Fl             []*FirstLevel
}

func (r *Root) GetId() string {
	return fmt.Sprintf("Root=%s", *r.Id)
}

type FirstLevel struct {
	Id             *string
	ParentInstance GetIdInterface
	Sl             []*SecondLevel
}

func (f *FirstLevel) GetId() string {
	return fmt.Sprintf("%s,FirstLevel=%s", f.ParentInstance.GetId(), *f.Id)
}

type SecondLevel struct {
	Id             *string
	ParentInstance GetIdInterface
}

func (s *SecondLevel) GetId() string {
	return fmt.Sprintf("%s,SecondLevel=%s", s.ParentInstance.GetId(), *s.Id)
}

func GetIdWorker() {
	rId := "r0"
	r := Root{
		Id:             &rId,
		ParentInstance: nil,
	}

	flId := "fl1"
	fl := FirstLevel{
		Id:             &flId,
		ParentInstance: &r,
	}

	slId := "sl2"
	sl := SecondLevel{
		Id:             &slId,
		ParentInstance: &fl,
	}

	fmt.Println(sl.GetId())
}
