package api
import (
	"github.com/sjkyspa/stacks/controller/api/util"
)

type AppParams struct {
	Stack      string `json:"stackId"`
	Name       string `json:"name"`
	NeedDeploy bool `json:"needDeploy"`
}

type AppRouteParams struct {
	Route string `json:"route"`
}

type UpdateStackParams struct {
	Stack string `json:"stack"`
}

type App interface {
	Id() string
	Links() Links
	GetBuilds() (Builds, error)
	GetRoutes() (AppRoutes, error)
	GetBuild(id string) (Build, error)
	GetBuildByURI(uri string) (Build, error)
	GetStack() (Stack, error)
	GetEnvs() (map[string]string)
	SetEnv(envs map[string]interface{}) (error)
	UnsetEnv(keys []string) (error)
	CreateBuild(buildParams BuildParams) (Build, error)
	BindWithRoute(params AppRouteParams) (error)
	UnbindRoute(routeId string) (error)
	SwitchStack(params UpdateStackParams) (error)
	GetLogForTests(buildId, logType string, lines int, startTimeStamp string) (string, error)
}

type AppModel struct {
	ID              string     `json:"name"`
	Envs            map[string]string `json:"envs"`
	LinksArray      []Link     `json:"links"`
	BuildMapper     BuildMapper
	AppMapper       AppRepository
	StackRepository StackRepository
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a AppModel) GetEnvs() map[string]string {
	return a.Envs
}

func (a AppModel) SetEnv(envs map[string]interface{}) (err error) {
	err = a.AppMapper.SetEnv(a, envs)
	return
}

func (a AppModel) UnsetEnv(keys []string) (err error) {
	err = a.AppMapper.UnsetEnv(a, keys)
	return
}

func (a AppModel) Id() string {
	return a.ID
}

func (a AppModel) Links() Links {
	return LinksModel{
		Links: a.LinksArray,
	}
}

func (a AppModel) GetBuilds() (builds Builds, apiError error) {
	return a.BuildMapper.GetBuilds(a)
}

func (a AppModel) GetBuildByURI(uri string) (build Build, apiError error) {
	id, err := util.IDFromURI(uri)
	if err != nil {
		apiError = err
		return
	}
	return a.BuildMapper.GetBuild(a, id)
}

func (a AppModel) GetBuild(id string) (build Build, apiError error) {
	return a.BuildMapper.GetBuild(a, id)
}

func (a AppModel) CreateBuild(buildParams BuildParams) (build Build, apiErr error) {
	return a.BuildMapper.Create(a, buildParams)
}

func (a AppModel) GetStack() (stack Stack, apiErr error) {
	stackLink, err := a.Links().Link("stack")
	if err != nil {
		apiErr = err
		return
	}

	return a.StackRepository.GetStackByURI(stackLink.URI)
}

func (a AppModel) BindWithRoute(params AppRouteParams) (error) {
	return a.AppMapper.BindWithRoute(a, params)
}

func (a AppModel) UnbindRoute(routeId string) (error) {
	return a.AppMapper.UnbindRoute(a, routeId)
}

func (a AppModel) GetRoutes() (AppRoutes, error) {
	return a.AppMapper.GetRoutes(a)
}

func (a AppModel) SwitchStack(params UpdateStackParams) (error) {
	return a.AppMapper.SwitchStack(a.ID, params)
}

func (a AppModel) GetLogForTests(buildId, logType string, lines int, startTimeStamp string) (string, error) {
	return a.AppMapper.GetLog(a.ID, buildId, logType, lines, startTimeStamp)
}

type AppRef interface {
	Id() string
	Links() Links
}

type AppRefModel struct {
	IDField     string     `json:"name"`
	LinksField  []Link     `json:"links"`
	BuildMapper BuildMapper
}

func (arm AppRefModel) Id() string {
	return arm.IDField
}

func (arm AppRefModel) Links() Links {
	return LinksModel{
		Links: arm.LinksField,
	}
}

type Apps interface {
	Count() int
	First() Apps
	Last() Apps
	Prev() Apps
	Next() Apps
	Items() []AppRef
}

type AppsModel struct {
	CountField int            `json:"count"`
	SelfField  string         `json:"self"`
	FirstField string         `json:"first"`
	LastField  string         `json:"last"`
	PrevField  string         `json:"prev"`
	NextField  string         `json:"next"`
	ItemsField []AppRefModel  `json:"items"`
	AppMapper  AppRepository
}

func (apps AppsModel) Count() int {
	return apps.CountField
}
func (apps AppsModel) Self() Apps {
	return nil
}
func (apps AppsModel) First() Apps {
	return nil
}
func (apps AppsModel) Last() Apps {
	return nil
}
func (apps AppsModel) Prev() Apps {
	return nil
}
func (apps AppsModel) Next() Apps {
	return nil
}

func (apps AppsModel) Items() []AppRef {
	items := make([]AppRef, 0)
	for _, app := range apps.ItemsField {
		items = append(items, app)
	}
	return items
}

type AppRouteModel struct {
	IDField      string     `json:"id"`
	PathField    string        `json:"path"`
	DomainField  SimpleDomain        `json:"domain"`
	CreatedField string        `json:"created"`
	LinksArray   []Link     `json:"links"`
}

type AppRoutes interface {
	Count() int
	First() (routes AppRoutes, apiError error)
	Last() (routes AppRoutes, apiError error)
	Prev() (routes AppRoutes, apiError error)
	Next() (routes AppRoutes, apiError error)
	Items() []AppRouteModel
}

type AppRoutesModel struct {
	CountField int            `json:"count"`
	SelfField  string         `json:"self"`
	FirstField string         `json:"first"`
	LastField  string         `json:"last"`
	PrevField  string         `json:"prev"`
	NextField  string         `json:"next"`
	ItemsField []AppRouteModel  `json:"items"`
	AppRepo    AppRepository
}

func (appRoutes AppRoutesModel) Count() int {
	return appRoutes.CountField
}
func (appRoutes AppRoutesModel) Self() (routes AppRoutes, apiError error) {
	if "" == appRoutes.SelfField {
		return
	}

	routes, apiError = appRoutes.AppRepo.GetRoutesByURI(appRoutes.SelfField)
	return
}
func (appRoutes AppRoutesModel) First() (routes AppRoutes, apiError error) {
	if "" == appRoutes.FirstField {
		return
	}

	routes, apiError = appRoutes.AppRepo.GetRoutesByURI(appRoutes.FirstField)
	return
}
func (appRoutes AppRoutesModel) Last() (routes AppRoutes, apiError error) {
	if "" == appRoutes.LastField {
		return
	}

	routes, apiError = appRoutes.AppRepo.GetRoutesByURI(appRoutes.LastField)
	return
}
func (appRoutes AppRoutesModel) Prev() (routes AppRoutes, apiError error) {
	if "" == appRoutes.PrevField {
		return
	}

	routes, apiError = appRoutes.AppRepo.GetRoutesByURI(appRoutes.PrevField)
	return
}
func (appRoutes AppRoutesModel) Next() (routes AppRoutes, apiError error) {
	if "" == appRoutes.NextField {
		return
	}

	routes, apiError = appRoutes.AppRepo.GetRoutesByURI(appRoutes.NextField)
	return
}

func (appRoutes AppRoutesModel) Items() []AppRouteModel {
	items := make([]AppRouteModel, 0)
	for _, app := range appRoutes.ItemsField {
		items = append(items, app)
	}
	return items
}
