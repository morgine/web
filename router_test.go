package web_test

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/morgine/web"
//	"testing"
//)
//
//func TestNewEngine(t *testing.T) {
//	engine := web.NewEngine()
//	type register struct {
//		group    *web.group
//		Filters  []*mockFilterHandler
//		Handlers []*mockHandler
//	}
//	type groupCase struct {
//		register    *register
//		groupRoutes []*web.route
//	}
//	type routeCase struct {
//		Method
//	}
//	groupCases := []*groupCase{
//		{
//			register: &register{
//				group: &web.group{
//					ID:      1,
//					Name:    "Admin Accounts",
//					Comment: "provides admin accounts API",
//					Path:    "/admin",
//				},
//				Filters: []*mockFilterHandler{
//					{"Auth Filter"},
//				},
//				Handlers: []*mockHandler{
//					{
//						"DELETE",
//						"/picture",
//						"Remove Picture Handler",
//						[]*mockFilterHandler{
//							{"Admin Filter"},
//						},
//					},
//				},
//			},
//			groupRoutes: []*web.route{
//				{
//					Method: "DELETE",
//					Path: "/admin/picture",
//					Handlers: []web.Handler{
//						&mockFilterHandler{"Auth Filter"},
//						&mockFilterHandler{"Admin Filter"},
//						&mockHandler {
//							"DELETE",
//							"/picture",
//							"Remove Picture Handler",
//							nil,
//						},
//					},
//				},
//			},
//		},
//		{
//			register: &register{
//				group: &web.group{
//					ID:      2,
//					Name:    "User Accounts",
//					Comment: "provides user accounts API",
//					Path:    "/user",
//				},
//				Filters: []*mockFilterHandler{
//					{"Auth Filter"},
//				},
//				Handlers: []*mockHandler{
//					{
//						"POST",
//						"/avatar",
//						"Upload Avatar Handler",
//						[]*mockFilterHandler{
//							{"User Filter"},
//						},
//					},
//					{
//						"POST",
//						"/reset",
//						"Reset Password Handler",
//						nil,
//					},
//				},
//			},
//			groupRoutes: []*web.route{
//				{
//					Method: "POST",
//					Path: "/user/avatar",
//					Handlers: []web.Handler {
//						&mockFilterHandler{"Auth Filter"},
//						&mockFilterHandler{"Self Filter"},
//						&mockHandler {
//							"POST",
//							"/avatar",
//							"Upload Avatar Handler",
//							nil,
//						},
//					},
//				},
//				{
//					Method: "POST",
//					Path: "/user/reset",
//					Handlers: []web.Handler {
//						&mockFilterHandler{"Auth Filter"},
//						&mockHandler {
//							"POST",
//							"/reset",
//							"Reset Password Handler",
//							nil,
//						},
//					},
//				},
//			},
//		},
//	}
//
//	var needGroups = []*web.group{
//		{
//			ID:      1,
//			Name:    "Admin Accounts",
//			Comment: "provides admin accounts API",
//			Path:    "/admin",
//		},
//		{
//			ID:      2,
//			Name:    "User Accounts",
//			Comment: "provides user accounts API",
//			Path:    "/user",
//		},
//	}
//
//	var needFilters = []*mockFilterHandler{
//		{"Auth Filter"},
//		{"Admin Filter"},
//		{"User Filter"},
//	}
//
//	// 添加 groups, Filters, Handlers
//	for _, c := range groupCases {
//		r := c.register
//		router := engine.group(r.group.Name, r.group.Comment, r.group.Path)
//		for _, mf := range r.Filters {
//			router = router.Use(mf)
//		}
//		for _, mh := range r.Handlers {
//			pr := router
//			for _, filter := range mh.PrivateFilters {
//				pr = pr.Use(filter)
//			}
//			pr.Handle(mh.Method, mh.Path, mh)
//		}
//	}
//
//	// groups, Filters, routes 容器
//	container := engine.Container()
//
//	var gotFilters = container.Filters()
//	if need, got := jsonStr(needFilters), jsonStr(gotFilters); need != got {
//		t.Fatalf("need: %v, \ngot: %v\n", need, got)
//	}
//
//	// 容器中的 groups
//	getGroups := container.Groups()
//	if need, got := jsonStr(needGroups), jsonStr(getGroups); need != got {
//		t.Fatalf("need: %v, \ngot: %v\n", need, got)
//	}
//
//	// group 中的 routes
//	for idx, group := range getGroups {
//		gotRoutes := container.Routes(group.ID)
//		needRoutes := groupCases[idx].groupRoutes
//		if need, got := jsonStr(needRoutes), jsonStr(gotRoutes); need != got {
//			t.Fatalf("need: %v, \ngot: %v\n", need, got)
//		}
//	}
//
//	// 路由 matcher
//	matcher := engine.Matcher()
//}
//
//func jsonStr(v interface{}) string {
//	data, err := json.Marshal(v)
//	if err != nil {
//		panic(err)
//	}
//	return string(data)
//}
//
//type mockFilterHandler struct {
//	Name string
//}
//
//func (j *mockFilterHandler) Handle(ctx *web.Context) error {
//	return nil
//}
//
//type mockHandler struct {
//	Method, Path, Name string
//	PrivateFilters            []*mockFilterHandler
//}
//
//func (m *mockHandler) Handle(ctx *web.Context) error {
//	return nil
//}
