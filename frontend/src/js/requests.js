
import axios from 'axios'

axios.interceptors.request.use(
  config => {
    config.data = JSON.stringify(config.data);
    config.headers = {
      // 'Content-Type': 'application/x-www-form-urlencoded'
      'Content-Type': 'multipart/form-data'
    };
    return config;
  },
  err => {
    return Promise.reject(err);
  }
);

var HOST = ''

export function get (url) {
  return new Promise(function (resolve, reject) {
    axios.get(url)
    .then(res => { resolve(res) })
    .catch(err => {
      if (err.response.status == 403) {
        reject(err)
      }
    })
  })
}

export function post (url, data) {
  return new Promise(function (resolve, reject) {
    axios.post(url, data)
    .then(res => { resolve(res) })
    .catch(err => {
      if (err.response.status == 403) {
        reject(err)
      }
    })
  })
}

if (process.env.NODE_ENV === 'development') {
  axios.defaults.baseURL = 'http://localhost:8000'
}

export function getBlogs() {
  return get(HOST + '/api/blogs')
}

export function getBlog(id, params) {
  return get(HOST + '/api/blogs/' + id, params)
}

export function getHtml(path) {
  // return get(HOST + '/statics/' + path)
  console.log(path)
  return get("http://localhost/statics/interface.md")
}
