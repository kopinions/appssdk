package api

type Image struct {
	Name string `json:"image"`
	Mem  int `json:"mem"`
	Cpus float64 `json:"cpus"`
}

type ServiceDefinition struct {
	Build     Image `json:"build"`
	Verify    Image `json:"verify"`
	Env       map[string]string `json:"environment"`
	Links     []string `json:"links"`
	Health    []HealthCheck `json:"health"`
	Volumes   []Volume `json:"volumes"`
	Exposes   int `json:"expose"`
	Image     string `json:"image"`
	cpu       float64 `json:"cpus"`
	mem       float64 `json:"mem"`
	instances int `json:"instances"`
	name      string `json:"name"`
}

func (sd ServiceDefinition) GetLinks() []string {
	return sd.Links
}

func (sd ServiceDefinition) GetEnv() map[string]string {
	return sd.Env
}

func (sd ServiceDefinition) GetVolumes() []Volume {
	return sd.Volumes
}

func (sd ServiceDefinition) GetHealthChecks() []HealthCheck {
	return sd.Health
}

func (sd ServiceDefinition) GetExpose() []int {
	exposes := make([]int, 1)
	exposes[0] = sd.Exposes
	return exposes
}

func (sd ServiceDefinition) GetImage() string {
	return sd.Image
}

func (sd ServiceDefinition) CPU() float64 {
	return sd.cpu
}

func (sd ServiceDefinition) Mem() float64 {
	return sd.mem
}

func (sd ServiceDefinition) Instances() int {
	return sd.instances
}

func (sd ServiceDefinition) Name() string {
	return sd.name
}

func (sd ServiceDefinition) IsBuildable() bool {
	return sd.Build.Name != ""
}

func (sd ServiceDefinition) GetBuild() Image {
	return sd.Build
}

func (sd ServiceDefinition) GetVerify() Image {
	return sd.Verify
}

type Template struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type Volume struct {
	ContainerPath string `json:"container"`
	HostPath      string `json:"host"`
	Mode          string `json:"mode"`
}

type HealthCheck struct {

}

type Service interface {
	GetEnv() map[string]string
	GetImage() string
	CPU() float64
	Mem() float64
	Instances() int
	Name() string
	GetVolumes() []Volume
	GetLinks() []string
	GetHealthChecks() []HealthCheck
	GetExpose() []int
	IsBuildable() bool
	GetBuild() Image
	GetVerify() Image
}

type Stack interface {
	Id() string
	Name() string
	Type() string
	Links() Links
	GetBuildImage() Image
	GetVerifyImage() Image
	GetTemplateCode() string
	GetServices() map[string]Service
	Update(stackDefinition map[string]interface{}) error
	Publish() error
	UnPublish() error
}

type StackModel struct {
	IDField     string `json:"id"`
	NameField   string `json:"name"`
	LinksField  []Link `json:"links"`
	TypeField   string `json:"type"`
	Services    map[string]ServiceDefinition `json:"services"`
	Template    Template `json:"template"`
	StackMapper StackRepository
}

func (a StackModel) Id() string {
	return a.IDField
}

func (a StackModel) Name() string {
	return a.NameField
}

func (s StackModel) Type() string {
	return s.TypeField
}

func (a StackModel) Links() Links {
	return LinksModel{
		Links: a.LinksField,
	}
}

func (a StackModel) GetBuildImage() Image {
	for _, v := range a.Services {
		if v.Build != (Image{}) {
			return v.Build
		}
	}
	return Image{}
}

func (a StackModel) GetVerifyImage() Image {
	for _, v := range a.Services {
		if v.Verify != (Image{}) {
			return v.Verify
		}
	}
	return Image{}
}

func (a StackModel) GetTemplateCode() string {
	return a.Template.URI
}

func (s StackModel) Update(stackDefinition map[string]interface{}) (err error) {
	err = s.StackMapper.Update(s.Id(), stackDefinition)
	return
}

func (s StackModel) Publish() (err error) {
	err = s.StackMapper.Publish(s.Id())
	return
}

func (s StackModel) UnPublish() (err error) {
	err = s.StackMapper.UnPublish(s.Id())
	return
}

func (s StackModel) GetServices() map[string]Service {
	services := make(map[string]Service, 1)

	for key, value := range s.Services {
		services[key] = value
	}
	return services
}

type Stacks interface {
	Count() int
	First() Stacks
	Last() Stacks
	Prev() Stacks
	Next() Stacks
	Items() []Stack
}

type StacksModel struct {
	CountField int            `json:"count"`
	SelfField  string         `json:"self"`
	FirstField string         `json:"first"`
	LastField  string         `json:"last"`
	PrevField  string         `json:"prev"`
	NextField  string         `json:"next"`
	ItemsField []StackModel   `json:"items"`
}

func (stacks StacksModel) Count() int {
	return stacks.CountField
}

func (stacks StacksModel) Self() Stacks {
	return nil
}
func (stacks StacksModel) First() Stacks {
	return nil
}
func (stacks StacksModel) Last() Stacks {
	return nil
}
func (stacks StacksModel) Prev() Stacks {
	return nil
}
func (stacks StacksModel) Next() Stacks {
	return nil
}

func (stacks StacksModel) Items() []Stack {
	items := make([]Stack, 0)
	for _, stack := range stacks.ItemsField {
		items = append(items, stack)
	}
	return items
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
