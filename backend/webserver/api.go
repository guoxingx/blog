package webserver

import (
	"encoding/json"
	"net/http"

	hw "backend/httpwrap"
	"backend/models"
	"backend/utils"
)

func handleIndex(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	if r.Method != "GET" {
		return 405, nil
	}

	return 200, hw.NewRestSuccess(nil)
}

func handleBlogs(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	if r.Method == "GET" {
		logger.Info("get blogs")
		logger.Debugf("url: %s", r.URL.RequestURI())

		blogs, err := models.ListBlogs(db)
		if err != nil {
			logger.Warnf("failed to list blogs: %s", err)
			return 200, hw.NewRestError(codeLoadDataFailed, "load data failed")
		}

		res := make([]*blog, len(blogs))
		for i, m := range blogs {
			res[i] = (&blog{}).fromModel(m)
		}

		logger.Warnf("request blogs success with %d blogs", len(res))
		return 200, hw.NewRestSuccess(res)

	} else if r.Method == "POST" {
		logger.Info("post blogs")

		var params blog
		err := json.NewDecoder(r.Body).Decode(&params)
		logger.Debugf("url: %s, params: %v", r.URL.RequestURI(), params)
		if err != nil {
			logger.Warnf("failed to load params: %v", err)
			return 200, hw.NewRestError(codeInvalidParams, "invalid params")
		}

		m, err := params.toModel()
		if err != nil {
			logger.Warnf("failed to parse params: %v", err)
			return 200, hw.NewRestError(codeInvalidParams, "invalid params")
		}
		err = m.Save(db)
		if err != nil {
			logger.Errorf("failed to save blog: %v", err)
			return 200, hw.NewRestError(codeProcessFailed, "invalid params")
		}

		params.fromModel(m)
		logger.Debugf("post blog success with response %v", params)
		return 200, hw.NewRestSuccess(params)

	}

	logger.Warnf("invalid method %s for request blogs", r.Method)
	return 405, nil
}

func handleBlog(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	if r.Method != "GET" {
		logger.Warnf("invalid method %s for request blog", r.Method)
		return 405, nil
	}
	logger.Info("get blog")
	logger.Debugf("url: %s", r.URL.RequestURI())

	id := utils.ToInt(hw.GetURLLastPart(r))
	if id == 0 {
		logger.Warnf("invalid blog id %d", id)
		return 200, hw.NewRestError(codeInvalidParams, "invalid id")
	}
	m := models.Blog{ID: id}
	err := m.Load(db)
	if err != nil {
		logger.Warnf("failed to find blog with id %d: %s", id, err)
		return 200, hw.NewRestError(codeInvalidParams, "invalid id")
	}

	res := (&blog{}).fromModel(&m)
	logger.Debugf("request blog success with response %v", res)
	return 200, hw.NewRestSuccess(res)
}
