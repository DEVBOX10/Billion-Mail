"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["885"],{6768:function(e,t,n){function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter(function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable})),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach(function(t){var r;r=n[t],t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach(function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))})}return e}function i(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=Array(t);n<t;n++)r[n]=e[n];return r}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter(function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable})),n.push.apply(n,r)}return n}function c(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach(function(t){var r;r=n[t],t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach(function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))})}return e}function l(e){return function t(){for(var n=this,r=arguments.length,o=Array(r),i=0;i<r;i++)o[i]=arguments[i];return o.length>=e.length?e.apply(this,o):function(){for(var e=arguments.length,r=Array(e),i=0;i<e;i++)r[i]=arguments[i];return t.apply(n,[].concat(o,r))}}}function s(e){return({}).toString.call(e).includes("Object")}function u(e){return"function"==typeof e}n.d(t,{Z:()=>A});var d,f,p=l(function(e,t){throw Error(e[t]||e.default)})({initialIsRequired:"initial state is required",initialType:"initial state should be an object",initialContent:"initial state shouldn't be an empty object",handlerType:"handler should be an object or a function",handlersType:"all handlers should be a functions",selectorType:"selector should be a function",changeType:"provided value of changes should be an object",changeField:'it seams you want to change a field in the state which is not specified in the "initial" state',default:"an unknown error accured in `state-local` package"}),h=function(e,t){return s(t)||p("changeType"),Object.keys(t).some(function(t){return!Object.prototype.hasOwnProperty.call(e,t)})&&p("changeField"),t},v=function(e){u(e)||p("selectorType")},m=function(e){u(e)||s(e)||p("handlerType"),s(e)&&Object.values(e).some(function(e){return!u(e)})&&p("handlersType")},b=function(e){e||p("initialIsRequired"),s(e)||p("initialType"),!Object.keys(e).length&&p("initialContent")};function g(e,t){return u(t)?t(e.current):t}function y(e,t){return e.current=c(c({},e.current),t),t}function j(e,t,n){return u(t)?t(e.current):Object.keys(n).forEach(function(n){var r;return null===(r=t[n])||void 0===r?void 0:r.call(t,e.current[n])}),n}var O={configIsRequired:"the configuration object is required",configType:"the configuration object should be an object",default:"an unknown error accured in `@monaco-editor/loader` package",deprecation:"Deprecation warning!\n    You are using deprecated way of configuration.\n\n    Instead of using\n      monaco.config({ urls: { monacoBase: '...' } })\n    use\n      monaco.config({ paths: { vs: '...' } })\n\n    For more please check the link https://github.com/suren-atoyan/monaco-loader#config\n  "},w=(d=function(e,t){throw Error(e[t]||e.default)},function e(){for(var t=this,n=arguments.length,r=Array(n),o=0;o<n;o++)r[o]=arguments[o];return r.length>=d.length?d.apply(this,r):function(){for(var n=arguments.length,o=Array(n),i=0;i<n;i++)o[i]=arguments[i];return e.apply(t,[].concat(r,o))}})(O);let k=function(e){return e||w("configIsRequired"),!({}).toString.call(e).includes("Object")&&w("configType"),e.urls?(console.warn(O.deprecation),{paths:{vs:e.urls.monacoBase}}):e},x=function(){for(var e=arguments.length,t=Array(e),n=0;n<e;n++)t[n]=arguments[n];return function(e){return t.reduceRight(function(e,t){return t(e)},e)}};var z={type:"cancelation",msg:"operation is manually canceled"};let S=function(e){var t=!1,n=new Promise(function(n,r){e.then(function(e){return t?r(z):n(e)}),e.catch(r)});return n.cancel=function(){return t=!0},n};var C=function(e){if(Array.isArray(e))return e}(f=({create:function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};b(e),m(t);var n={current:e},r=l(j)(n,t),o=l(y)(n),i=l(h)(e),a=l(g)(n);return[function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:function(e){return e};return v(e),e(n.current)},function(e){(function(){for(var e=arguments.length,t=Array(e),n=0;n<e;n++)t[n]=arguments[n];return function(e){return t.reduceRight(function(e,t){return t(e)},e)}})(r,o,i,a)(e)}]}}).create({config:{paths:{vs:"https://cdn.jsdelivr.net/npm/monaco-editor@0.52.2/min/vs"}},isInitialized:!1,resolve:null,reject:null,monaco:null}))||function(e,t){if("undefined"!=typeof Symbol&&Symbol.iterator in Object(e)){var n=[],r=!0,o=!1,i=void 0;try{for(var a,c=e[Symbol.iterator]();!(r=(a=c.next()).done)&&(n.push(a.value),n.length!==t);r=!0);}catch(e){o=!0,i=e}finally{try{r||null==c.return||c.return()}finally{if(o)throw i}}return n}}(f,2)||function(e,t){if(e){if("string"==typeof e)return i(e,2);var n=Object.prototype.toString.call(e).slice(8,-1);if("Object"===n&&e.constructor&&(n=e.constructor.name),"Map"===n||"Set"===n)return Array.from(e);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return i(e,t)}}(f,2)||function(){throw TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}(),P=C[0],B=C[1];function T(e){return document.body.appendChild(e)}function E(e){var t,n,r=P(function(e){return{config:e.config,reject:e.reject}}),o=(t="".concat(r.config.paths.vs,"/loader.js"),n=document.createElement("script"),t&&(n.src=t),n);return o.onload=function(){return e()},o.onerror=r.reject,o}function I(){var e=P(function(e){return{config:e.config,resolve:e.resolve,reject:e.reject}}),t=window.require;t.config(e.config),t(["vs/editor/editor.main"],function(t){Z(t),e.resolve(t)},function(t){e.reject(t)})}function Z(e){P().monaco||B({monaco:e})}var $=new Promise(function(e,t){return B({resolve:e,reject:t})});let A={config:function(e){var t=k(e),n=t.monaco,r=function(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],!(t.indexOf(n)>=0)&&Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}(t,["monaco"]);B(function(e){return{config:function e(t,n){return Object.keys(n).forEach(function(r){n[r]instanceof Object&&t[r]&&Object.assign(n[r],e(t[r],n[r]))}),o(o({},t),n)}(e.config,r),monaco:n}})},init:function(){var e=P(function(e){return{monaco:e.monaco,isInitialized:e.isInitialized,resolve:e.resolve}});if(!e.isInitialized){if(B({isInitialized:!0}),e.monaco)return e.resolve(e.monaco),S($);if(window.monaco&&window.monaco.editor)return Z(window.monaco),e.resolve(window.monaco),S($);x(T,E)(I)}return S($)},__getMonacoInstance:function(){return P(function(e){return e.monaco})}}},4944:function(e,t,n){n.d(t,{L:()=>r});function r(e,t){if(!e)return;let n=document.createElement("a");n.href=e,void 0!==t&&(n.download=t),document.body.appendChild(n),n.click(),document.body.removeChild(n)}},3012:function(e,t,n){n.d(t,{W3:()=>p,ZP:()=>v});var r=n(209),o=n(1321),i=n(4124),a=n(6169),c=n(1579),l=n(8755);let s={fontWeightActive:"400"},u={name:"Breadcrumb",common:l.Z,self:function(e){let{fontSize:t,textColor3:n,textColor2:r,borderRadius:o,buttonColor2Hover:i,buttonColor2Pressed:a}=e;return Object.assign(Object.assign({},s),{fontSize:t,itemLineHeight:"1.25",itemTextColor:n,itemTextColorHover:r,itemTextColorPressed:r,itemTextColorActive:r,itemBorderRadius:o,itemColorHover:i,itemColorPressed:a,separatorColor:n})}};var d=n(2249);let f=(0,d.cB)("breadcrumb",`
 white-space: nowrap;
 cursor: default;
 line-height: var(--n-item-line-height);
`,[(0,d.c)("ul",`
 list-style: none;
 padding: 0;
 margin: 0;
 `),(0,d.c)("a",`
 color: inherit;
 text-decoration: inherit;
 `),(0,d.cB)("breadcrumb-item",`
 font-size: var(--n-font-size);
 transition: color .3s var(--n-bezier);
 display: inline-flex;
 align-items: center;
 `,[(0,d.cB)("icon",`
 font-size: 18px;
 vertical-align: -.2em;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `),(0,d.c)("&:not(:last-child)",[(0,d.cM)("clickable",[(0,d.cE)("link",`
 cursor: pointer;
 `,[(0,d.c)("&:hover",`
 background-color: var(--n-item-color-hover);
 `),(0,d.c)("&:active",`
 background-color: var(--n-item-color-pressed); 
 `)])])]),(0,d.cE)("link",`
 padding: 4px;
 border-radius: var(--n-item-border-radius);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 position: relative;
 `,[(0,d.c)("&:hover",`
 color: var(--n-item-text-color-hover);
 `,[(0,d.cB)("icon",`
 color: var(--n-item-text-color-hover);
 `)]),(0,d.c)("&:active",`
 color: var(--n-item-text-color-pressed);
 `,[(0,d.cB)("icon",`
 color: var(--n-item-text-color-pressed);
 `)])]),(0,d.cE)("separator",`
 margin: 0 8px;
 color: var(--n-separator-color);
 transition: color .3s var(--n-bezier);
 user-select: none;
 -webkit-user-select: none;
 `),(0,d.c)("&:last-child",[(0,d.cE)("link",`
 font-weight: var(--n-font-weight-active);
 cursor: unset;
 color: var(--n-item-text-color-active);
 `,[(0,d.cB)("icon",`
 color: var(--n-item-text-color-active);
 `)]),(0,d.cE)("separator",`
 display: none;
 `)])])]),p=(0,c.U)("n-breadcrumb"),h=Object.assign(Object.assign({},o.Z.props),{separator:{type:String,default:"/"}}),v=(0,r.aZ)({name:"Breadcrumb",props:h,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:n}=(0,i.ZP)(e),c=(0,o.Z)("Breadcrumb","-breadcrumb",f,u,e,t);(0,r.JJ)(p,{separatorRef:(0,r.Vh)(e,"separator"),mergedClsPrefixRef:t});let l=(0,r.Fl)(()=>{let{common:{cubicBezierEaseInOut:e},self:{separatorColor:t,itemTextColor:n,itemTextColorHover:r,itemTextColorPressed:o,itemTextColorActive:i,fontSize:a,fontWeightActive:l,itemBorderRadius:s,itemColorHover:u,itemColorPressed:d,itemLineHeight:f}}=c.value;return{"--n-font-size":a,"--n-bezier":e,"--n-item-text-color":n,"--n-item-text-color-hover":r,"--n-item-text-color-pressed":o,"--n-item-text-color-active":i,"--n-separator-color":t,"--n-item-color-hover":u,"--n-item-color-pressed":d,"--n-item-border-radius":s,"--n-font-weight-active":l,"--n-item-line-height":f}}),s=n?(0,a.F)("breadcrumb",void 0,l,e):void 0;return{mergedClsPrefix:t,cssVars:n?void 0:l,themeClass:null==s?void 0:s.themeClass,onRender:null==s?void 0:s.onRender}},render(){var e;return null===(e=this.onRender)||void 0===e||e.call(this),(0,r.h)("nav",{class:[`${this.mergedClsPrefix}-breadcrumb`,this.themeClass],style:this.cssVars,"aria-label":"Breadcrumb"},(0,r.h)("ul",null,this.$slots))}})},6666:function(e,t,n){n.d(t,{Z:()=>c});var r=n(209),o=n(8282),i=n(4934),a=n(3012);let c=(0,r.aZ)({name:"BreadcrumbItem",props:{separator:String,href:String,clickable:{type:Boolean,default:!0},onClick:Function},slots:Object,setup(e,{slots:t}){let n=(0,r.f3)(a.W3,null);if(!n)return()=>null;let{separatorRef:c,mergedClsPrefixRef:l}=n,s=function(e=i.j?window:null){let t=()=>{let{hash:t,host:n,hostname:r,href:o,origin:i,pathname:a,port:c,protocol:l,search:s}=(null==e?void 0:e.location)||{};return{hash:t,host:n,hostname:r,href:o,origin:i,pathname:a,port:c,protocol:l,search:s}},n=(0,r.iH)(t()),o=()=>{n.value=t()};return(0,r.bv)(()=>{e&&(e.addEventListener("popstate",o),e.addEventListener("hashchange",o))}),(0,r.SK)(()=>{e&&(e.removeEventListener("popstate",o),e.removeEventListener("hashchange",o))}),n}(),u=(0,r.Fl)(()=>e.href?"a":"span"),d=(0,r.Fl)(()=>s.value.href===e.href?"location":null);return()=>{let{value:n}=l;return(0,r.h)("li",{class:[`${n}-breadcrumb-item`,e.clickable&&`${n}-breadcrumb-item--clickable`]},(0,r.h)(u.value,{class:`${n}-breadcrumb-item__link`,"aria-current":d.value,href:e.href,onClick:e.onClick},t),(0,r.h)("span",{class:`${n}-breadcrumb-item__separator`,"aria-hidden":"true"},(0,o.gI)(t.separator,()=>{var t;return[null!==(t=e.separator)&&void 0!==t?t:c.value]})))}}})},8951:function(e,t,n){n.d(t,{Z:()=>m});var r=n(5083),o=n(2370),i=n(209),a=n(4131),c=n(1321),l=n(4124),s=n(6169),u=n(2249);let d={name:"Spin",common:n(8755).Z,self:function(e){let{opacityDisabled:t,heightTiny:n,heightSmall:r,heightMedium:o,heightLarge:i,heightHuge:a,primaryColor:c,fontSize:l}=e;return{fontSize:l,textColor:c,sizeTiny:n,sizeSmall:r,sizeMedium:o,sizeLarge:i,sizeHuge:a,color:c,opacitySpinning:t}}};var f=n(3291);let p=(0,u.c)([(0,u.c)("@keyframes spin-rotate",`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),(0,u.cB)("spin-container",`
 position: relative;
 `,[(0,u.cB)("spin-body",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[(0,f.h)()])]),(0,u.cB)("spin-body",`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),(0,u.cB)("spin",`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[(0,u.cM)("rotate",`
 animation: spin-rotate 2s linear infinite;
 `)]),(0,u.cB)("spin-description",`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),(0,u.cB)("spin-content",`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[(0,u.cM)("spinning",`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]),h={small:20,medium:18,large:16},v=Object.assign(Object.assign({},c.Z.props),{contentClass:String,contentStyle:[Object,String],description:String,stroke:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},strokeWidth:Number,rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0},delay:Number}),m=(0,i.aZ)({name:"Spin",props:v,slots:Object,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:n}=(0,l.ZP)(e),a=(0,c.Z)("Spin","-spin",p,d,e,t),f=(0,i.Fl)(()=>{let{size:t}=e,{common:{cubicBezierEaseInOut:n},self:o}=a.value,{opacitySpinning:i,color:c,textColor:l}=o;return{"--n-bezier":n,"--n-opacity-spinning":i,"--n-size":"number"==typeof t?(0,r.BL)(t):o[(0,u.Tl)("size",t)],"--n-color":c,"--n-text-color":l}}),v=n?(0,s.F)("spin",(0,i.Fl)(()=>{let{size:t}=e;return"number"==typeof t?String(t):t[0]}),f,e):void 0,m=(0,o.Z)(e,["spinning","show"]),b=(0,i.iH)(!1);return(0,i.m0)(t=>{let n;if(m.value){let{delay:r}=e;if(r){n=window.setTimeout(()=>{b.value=!0},r),t(()=>{clearTimeout(n)});return}}b.value=m.value}),{mergedClsPrefix:t,active:b,mergedStrokeWidth:(0,i.Fl)(()=>{let{strokeWidth:t}=e;if(void 0!==t)return t;let{size:n}=e;return h["number"==typeof n?"medium":n]}),cssVars:n?void 0:f,themeClass:null==v?void 0:v.themeClass,onRender:null==v?void 0:v.onRender}},render(){var e,t;let{$slots:n,mergedClsPrefix:r,description:o}=this,c=n.icon&&this.rotate,l=(o||n.description)&&(0,i.h)("div",{class:`${r}-spin-description`},o||(null===(e=n.description)||void 0===e?void 0:e.call(n))),s=n.icon?(0,i.h)("div",{class:[`${r}-spin-body`,this.themeClass]},(0,i.h)("div",{class:[`${r}-spin`,c&&`${r}-spin--rotate`],style:n.default?"":this.cssVars},n.icon()),l):(0,i.h)("div",{class:[`${r}-spin-body`,this.themeClass]},(0,i.h)(a.Z,{clsPrefix:r,style:n.default?"":this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,class:`${r}-spin`}),l);return null===(t=this.onRender)||void 0===t||t.call(this),n.default?(0,i.h)("div",{class:[`${r}-spin-container`,this.themeClass],style:this.cssVars},(0,i.h)("div",{class:[`${r}-spin-content`,this.active&&`${r}-spin-content--spinning`,this.contentClass],style:this.contentStyle},n),(0,i.h)(i.uT,{name:"fade-in-transition"},{default:()=>this.active?s:null})):s}})}}]);