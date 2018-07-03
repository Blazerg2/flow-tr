package main

import (
	"flow-tr/controllers"
	"flow-tr/utils"
)

// AppRoutes define all routes in http application
func AppRoutes() utils.Routes {
	return utils.Routes{
		utils.Route{
			"Index",
			"GET",
			"/",
			controllers.Index,
		},
		utils.Route{
			"GetPageVar",
			"GET",
			"/pagevar/{pageVarId}",
			controllers.GetPageVar,
		},
		utils.Route{
			"GetPagesList",
			"GET",
			"/pages",
			controllers.GetPagesList,
		},
		utils.Route{
			"PostPagesList",
			"POST",
			"/pages",
			controllers.PostPagesList,
		},
	}
}
