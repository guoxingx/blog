(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-60c404b7"],{"11eb":function(t,n,e){"use strict";e.r(n);var o=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("div",{staticClass:"m-content-all"},[e("h2",[t._v("GX's Blog")]),e("h3",[t._v("Index")]),e("div",{staticClass:"blogs"},t._l(t.blogsShow,(function(t){return e("blog-cell",{key:t,attrs:{blog:t}})})),1),this.blogs.length>10?e("el-pagination",{staticClass:"page",attrs:{layout:"prev,pager,next",small:"",total:this.blogs.length,"page-size":10},on:{"current-change":t.changePage}}):t._e()],1)},r=[],c=(e("fb6a"),e("d3b7"),e("ec60")),a={name:"Blogs",components:{"blog-cell":function(){return e.e("chunk-42d24634").then(e.bind(null,"99d9"))}},data:function(){return{date:new Date,blogs:[],blogsShow:[]}},watch:{$route:function(t,n){console.log(t,n),this.listBlogs()}},created:function(){this.listBlogs()},methods:{listBlogs:function(){var t=this;Object(c["b"])().then((function(n){0==n.data.code&&n.data.data.length>0&&(t.blogs=n.data.data,t.blogsShow=t.blogs.slice(0,10))}))},changePage:function(t){this.blogsShow=this.blogs.slice(10*(t-1),10*t)}}},i=a,s=(e("bc4e"),e("2877")),l=Object(s["a"])(i,o,r,!1,null,null,null);n["default"]=l.exports},"1dde":function(t,n,e){var o=e("d039"),r=e("b622"),c=e("2d00"),a=r("species");t.exports=function(t){return c>=51||!o((function(){var n=[],e=n.constructor={};return e[a]=function(){return{foo:1}},1!==n[t](Boolean).foo}))}},8418:function(t,n,e){"use strict";var o=e("c04e"),r=e("9bf2"),c=e("5c6c");t.exports=function(t,n,e){var a=o(n);a in t?r.f(t,a,c(0,e)):t[a]=e}},ae40:function(t,n,e){var o=e("83ab"),r=e("d039"),c=e("5135"),a=Object.defineProperty,i={},s=function(t){throw t};t.exports=function(t,n){if(c(i,t))return i[t];n||(n={});var e=[][t],l=!!c(n,"ACCESSORS")&&n.ACCESSORS,u=c(n,0)?n[0]:s,f=c(n,1)?n[1]:void 0;return i[t]=!!e&&!r((function(){if(l&&!o)return!0;var t={length:-1};l?a(t,1,{enumerable:!0,get:s}):t[1]=1,e.call(t,u,f)}))}},bc4e:function(t,n,e){"use strict";var o=e("f444"),r=e.n(o);r.a},e8b5:function(t,n,e){var o=e("c6b6");t.exports=Array.isArray||function(t){return"Array"==o(t)}},f444:function(t,n,e){},fb6a:function(t,n,e){"use strict";var o=e("23e7"),r=e("861d"),c=e("e8b5"),a=e("23cb"),i=e("50c4"),s=e("fc6a"),l=e("8418"),u=e("b622"),f=e("1dde"),b=e("ae40"),g=f("slice"),d=b("slice",{ACCESSORS:!0,0:0,1:2}),h=u("species"),v=[].slice,p=Math.max;o({target:"Array",proto:!0,forced:!g||!d},{slice:function(t,n){var e,o,u,f=s(this),b=i(f.length),g=a(t,b),d=a(void 0===n?b:n,b);if(c(f)&&(e=f.constructor,"function"!=typeof e||e!==Array&&!c(e.prototype)?r(e)&&(e=e[h],null===e&&(e=void 0)):e=void 0,e===Array||void 0===e))return v.call(f,g,d);for(o=new(void 0===e?Array:e)(p(d-g,0)),u=0;g<d;g++,u++)g in f&&l(o,u,f[g]);return o.length=u,o}})}}]);
//# sourceMappingURL=chunk-60c404b7.aba8135e.js.map