(window.webpackJsonp=window.webpackJsonp||[]).push([[10],{78:function(t,e,n){"use strict";n.r(e),n.d(e,"frontMatter",(function(){return i})),n.d(e,"metadata",(function(){return c})),n.d(e,"toc",(function(){return s})),n.d(e,"default",(function(){return l}));var r=n(3),a=n(7),o=(n(0),n(84)),i={id:"annotations",title:"Annotations"},c={unversionedId:"annotations",id:"annotations",isDocsHomePage:!1,title:"Annotations",description:"Annotations let you extract data from a data source and use it to annotate a dashboard.",source:"@site/docs/annotations.md",slug:"/annotations",permalink:"/grafana-csv-datasource/annotations",editUrl:"https://github.com/marcusolsson/grafana-csv-datasource/edit/main/website/docs/annotations.md",version:"current",sidebar:"someSidebar",previous:{title:"Variables",permalink:"/grafana-csv-datasource/variables"}},s=[],u={toc:s};function l(t){var e=t.components,n=Object(a.a)(t,["components"]);return Object(o.b)("wrapper",Object(r.a)({},u,n,{components:e,mdxType:"MDXLayout"}),Object(o.b)("p",null,Object(o.b)("a",{parentName:"p",href:"https://grafana.com/docs/grafana/latest/dashboards/annotations"},"Annotations")," let you extract data from a data source and use it to annotate a dashboard."),Object(o.b)("p",null,"To use the CSV data source for annotations, follow the instructions on ",Object(o.b)("a",{parentName:"p",href:"https://grafana.com/docs/grafana/latest/dashboards/annotations/#querying-other-data-sources"},"Querying other data sources"),". Make sure to select the CSV from the list of data sources."),Object(o.b)("p",null,"Configure a query with ",Object(o.b)("em",{parentName:"p"},"at least")," two fields:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"A ",Object(o.b)("strong",{parentName:"li"},"String")," field for the annotation text"),Object(o.b)("li",{parentName:"ul"},"A ",Object(o.b)("strong",{parentName:"li"},"Time")," field for the annotation time")),Object(o.b)("p",null,"If you want to add titles or tags to the annotations, you can add additional ",Object(o.b)("strong",{parentName:"p"},"Fields")," with the appropriate types."),Object(o.b)("p",null,"For more information on how to configure a query, refer to ",Object(o.b)("a",{parentName:"p",href:"/grafana-csv-datasource/query-editor"},"Query editor"),"."))}l.isMDXComponent=!0},84:function(t,e,n){"use strict";n.d(e,"a",(function(){return p})),n.d(e,"b",(function(){return b}));var r=n(0),a=n.n(r);function o(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t}function i(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r)}return n}function c(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?i(Object(n),!0).forEach((function(e){o(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}function s(t,e){if(null==t)return{};var n,r,a=function(t,e){if(null==t)return{};var n,r,a={},o=Object.keys(t);for(r=0;r<o.length;r++)n=o[r],e.indexOf(n)>=0||(a[n]=t[n]);return a}(t,e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(t);for(r=0;r<o.length;r++)n=o[r],e.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(t,n)&&(a[n]=t[n])}return a}var u=a.a.createContext({}),l=function(t){var e=a.a.useContext(u),n=e;return t&&(n="function"==typeof t?t(e):c(c({},e),t)),n},p=function(t){var e=l(t.components);return a.a.createElement(u.Provider,{value:e},t.children)},f={inlineCode:"code",wrapper:function(t){var e=t.children;return a.a.createElement(a.a.Fragment,{},e)}},d=a.a.forwardRef((function(t,e){var n=t.components,r=t.mdxType,o=t.originalType,i=t.parentName,u=s(t,["components","mdxType","originalType","parentName"]),p=l(n),d=r,b=p["".concat(i,".").concat(d)]||p[d]||f[d]||o;return n?a.a.createElement(b,c(c({ref:e},u),{},{components:n})):a.a.createElement(b,c({ref:e},u))}));function b(t,e){var n=arguments,r=e&&e.mdxType;if("string"==typeof t||r){var o=n.length,i=new Array(o);i[0]=d;var c={};for(var s in e)hasOwnProperty.call(e,s)&&(c[s]=e[s]);c.originalType=t,c.mdxType="string"==typeof t?t:r,i[1]=c;for(var u=2;u<o;u++)i[u]=n[u];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,n)}d.displayName="MDXCreateElement"}}]);