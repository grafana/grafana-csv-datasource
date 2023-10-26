"use strict";(self.webpackChunkgrafana_csv_datasource_docs=self.webpackChunkgrafana_csv_datasource_docs||[]).push([[687],{2232:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>d,contentTitle:()=>i,default:()=>c,frontMatter:()=>r,metadata:()=>o,toc:()=>h});var a=s(5893),n=s(1151);const r={id:"query-editor",title:"Query editor"},i=void 0,o={id:"query-editor",title:"Query editor",description:"This page explains the what each part of the query editor does, and how you can configure it.",source:"@site/docs/query-editor.md",sourceDirName:".",slug:"/query-editor",permalink:"/grafana-csv-datasource/query-editor",draft:!1,unlisted:!1,editUrl:"https://github.com/grafana/grafana-csv-datasource/edit/main/website/docs/query-editor.md",tags:[],version:"current",frontMatter:{id:"query-editor",title:"Query editor"},sidebar:"someSidebar",previous:{title:"Configuration",permalink:"/grafana-csv-datasource/configuration"},next:{title:"Variables",permalink:"/grafana-csv-datasource/variables"}},d={},h=[{value:"Fields",id:"fields",level:3},{value:"Path",id:"path",level:3},{value:"HTTP",id:"http",level:4},{value:"Local",id:"local",level:4},{value:"Params",id:"params",level:3},{value:"Headers",id:"headers",level:3},{value:"Body",id:"body",level:3},{value:"Experimental",id:"experimental",level:3}];function l(e){const t=Object.assign({p:"p",h3:"h3",img:"img",strong:"strong",ul:"ul",li:"li",em:"em",h4:"h4",a:"a",code:"code",blockquote:"blockquote",admonition:"admonition"},(0,n.a)(),e.components);return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(t.p,{children:"This page explains the what each part of the query editor does, and how you can configure it."}),"\n",(0,a.jsx)(t.p,{children:"The query editor for the CSV data source consists of a number of tabs. Each tab configures a part of the query."}),"\n",(0,a.jsx)(t.h3,{id:"fields",children:"Fields"}),"\n",(0,a.jsx)(t.p,{children:(0,a.jsx)(t.img,{alt:"Fields",src:s(2397).Z+"",width:"1928",height:"380"})}),"\n",(0,a.jsxs)(t.p,{children:["The ",(0,a.jsx)(t.strong,{children:"Fields"})," tab is where you configure how to parse the data returned by the URL configured in the data source configuration."]}),"\n",(0,a.jsxs)(t.ul,{children:["\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Delimiter"})," defines how columns are separated in the file."]}),"\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Decimal separator"})," defines the character used to separate the integer and fractional part of a number."]}),"\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Skip leading rows"})," allows you to ignore a number of rows at the start of the data. This can be useful if your data contains comments, documentations, or other information before the data."]}),"\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Header"})," tells Grafana whether the first row contains the names of each column."]}),"\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Ignore unknown"})," toggles whether to display columns that aren't defined in the schema. This can be useful if you're only interested in a few columns."]}),"\n"]}),"\n",(0,a.jsxs)(t.p,{children:["By default, all columns in the CSV data are treated as text. If you want to parse a column into a specific type, such as Time or Number, you need to define a ",(0,a.jsx)(t.em,{children:"schema"}),"."]}),"\n",(0,a.jsxs)(t.ul,{children:["\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Field"}),' references a column in the CSV data. If no header is present, you can reference the columns by their order, for example "Field 1".']}),"\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Type"})," defines the type of the column. If the type is anything other than ",(0,a.jsx)(t.strong,{children:"String"}),", the data source tries to parse the data into the specified type. Any values that can't be parsed are ignored."]}),"\n"]}),"\n",(0,a.jsx)(t.h3,{id:"path",children:"Path"}),"\n",(0,a.jsx)(t.p,{children:"The contents of this tab depends on whether the data source is set to HTTP or Local mode. In both cases, the path is relative to the data source URL."}),"\n",(0,a.jsx)(t.h4,{id:"http",children:"HTTP"}),"\n",(0,a.jsx)(t.p,{children:(0,a.jsx)(t.img,{alt:"Path",src:s(7434).Z+"",width:"1932",height:"238"})}),"\n",(0,a.jsxs)(t.p,{children:["The drop-down box to the left lets you configure the ",(0,a.jsx)(t.strong,{children:"HTTP method"})," of the request sent to the URL and can be set to ",(0,a.jsx)(t.strong,{children:"GET"})," and ",(0,a.jsx)(t.strong,{children:"POST"}),"."]}),"\n",(0,a.jsxs)(t.p,{children:["The text box lets you append a path to the URL in the data source configuration. This can be used to dynamically change the request URL using ",(0,a.jsx)(t.a,{href:"https://grafana.com/docs/grafana/latest/variables/",children:"variables"}),"."]}),"\n",(0,a.jsxs)(t.p,{children:["For example, by setting the path to ",(0,a.jsx)(t.code,{children:"/movies/${movie}/summary"})," you can query the summary for any movie without having to change the query itself."]}),"\n",(0,a.jsx)(t.h4,{id:"local",children:"Local"}),"\n",(0,a.jsx)(t.p,{children:(0,a.jsx)(t.img,{alt:"Path",src:s(3538).Z+"",width:"1932",height:"232"})}),"\n",(0,a.jsxs)(t.p,{children:[(0,a.jsx)(t.strong,{children:"Relative path"})," lets you append a relative path to the one in the data source configuration. For example, you can use the same data source to load multiple files by setting the ",(0,a.jsx)(t.strong,{children:"Path"})," in the data source configuration to a directory, and then use the ",(0,a.jsx)(t.strong,{children:"Relative path"})," to load a file within that directory."]}),"\n",(0,a.jsx)(t.h3,{id:"params",children:"Params"}),"\n",(0,a.jsxs)(t.blockquote,{children:["\n",(0,a.jsx)(t.p,{children:"Only available in HTTP mode."}),"\n"]}),"\n",(0,a.jsx)(t.p,{children:(0,a.jsx)(t.img,{alt:"Params",src:s(4522).Z+"",width:"1932",height:"300"})}),"\n",(0,a.jsxs)(t.p,{children:["Add any parameters you'd like to send as part of the query string. For example, the parameters in the screenshot gets encoded as ",(0,a.jsx)(t.code,{children:"?category=movies"}),"."]}),"\n",(0,a.jsxs)(t.p,{children:["Both the ",(0,a.jsx)(t.strong,{children:"Key"})," and ",(0,a.jsx)(t.strong,{children:"Value"})," fields support ",(0,a.jsx)(t.a,{href:"https://grafana.com/docs/grafana/latest/variables/",children:"variables"}),"."]}),"\n",(0,a.jsx)(t.admonition,{type:"caution",children:(0,a.jsx)(t.p,{children:"Any query parameters that have been set by the administrator in the data source configuration has higher priority and overrides the parameters set by the query."})}),"\n",(0,a.jsx)(t.h3,{id:"headers",children:"Headers"}),"\n",(0,a.jsxs)(t.blockquote,{children:["\n",(0,a.jsx)(t.p,{children:"Only available in HTTP mode."}),"\n"]}),"\n",(0,a.jsx)(t.p,{children:(0,a.jsx)(t.img,{alt:"Headers",src:s(6741).Z+"",width:"1932",height:"300"})}),"\n",(0,a.jsx)(t.p,{children:"Add any parameters you'd like to send as HTTP headers."}),"\n",(0,a.jsxs)(t.p,{children:["Both the ",(0,a.jsx)(t.strong,{children:"Key"})," and ",(0,a.jsx)(t.strong,{children:"Value"})," fields support ",(0,a.jsx)(t.a,{href:"https://grafana.com/docs/grafana/latest/variables/",children:"variables"}),"."]}),"\n",(0,a.jsx)(t.h3,{id:"body",children:"Body"}),"\n",(0,a.jsxs)(t.blockquote,{children:["\n",(0,a.jsx)(t.p,{children:"Only available in HTTP mode."}),"\n"]}),"\n",(0,a.jsx)(t.p,{children:(0,a.jsx)(t.img,{alt:"Body",src:s(1988).Z+"",width:"1934",height:"642"})}),"\n",(0,a.jsx)(t.p,{children:"Sets the text to send as a request body."}),"\n",(0,a.jsxs)(t.ul,{children:["\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Syntax highlighting"})," sets the active syntax for the editor. This is only for visual purposes and doesn't change the actual request."]}),"\n"]}),"\n",(0,a.jsx)(t.admonition,{type:"info",children:(0,a.jsx)(t.p,{children:"Due to limitations in modern browsers, Grafana ignores the request body if the HTTP method is set to GET."})}),"\n",(0,a.jsx)(t.h3,{id:"experimental",children:"Experimental"}),"\n",(0,a.jsx)(t.p,{children:"Try out features that are currently in development. Each feature has a link in its tooltip that takes you to the feature request on GitHub where you can share your feedback."}),"\n",(0,a.jsx)(t.admonition,{type:"danger",children:(0,a.jsx)(t.p,{children:"Experimental features might be unstable and can be removed without notice."})}),"\n",(0,a.jsxs)(t.ul,{children:["\n",(0,a.jsxs)(t.li,{children:[(0,a.jsx)(t.strong,{children:"Enable regular expressions"})," lets you use regular expressions as field names in the schema. This lets you set the type for any field that matches the expression."]}),"\n"]})]})}const c=function(e={}){const{wrapper:t}=Object.assign({},(0,n.a)(),e.components);return t?(0,a.jsx)(t,Object.assign({},e,{children:(0,a.jsx)(l,e)})):l(e)}},1988:(e,t,s)=>{s.d(t,{Z:()=>a});const a=s.p+"assets/images/editor-body-6ea39c0e5233d24d162d82c634bbaea7.png"},2397:(e,t,s)=>{s.d(t,{Z:()=>a});const a=s.p+"assets/images/editor-fields-b8293a85e3500950b1109ccf3356fe11.png"},6741:(e,t,s)=>{s.d(t,{Z:()=>a});const a=s.p+"assets/images/editor-headers-8c732e5f73fc34c1a0005d2db9ced1ef.png"},3538:(e,t,s)=>{s.d(t,{Z:()=>a});const a=s.p+"assets/images/editor-local-path-d98216323e20deca2338383425b4dad3.png"},4522:(e,t,s)=>{s.d(t,{Z:()=>a});const a=s.p+"assets/images/editor-params-f7f56d3cab21e946da8bca23845e45f2.png"},7434:(e,t,s)=>{s.d(t,{Z:()=>a});const a=s.p+"assets/images/editor-path-1f1cec2bf001472dac78f5c6091b58a9.png"},1151:(e,t,s)=>{s.d(t,{a:()=>i});var a=s(7294);const n={},r=a.createContext(n);function i(e){const t=a.useContext(r);return a.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}}}]);