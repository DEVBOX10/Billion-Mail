"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["840"],{7610:function(e,o,t){t.d(o,{mj:()=>p,ZP:()=>g});var r=t(209),l=t(2121),i=t(1321),n=t(4124),a=t(6169),c=t(1579),d=t(5568),s=t(5291),v=t(2988),h=t(2249);let u=(0,h.cB)("layout",`
 color: var(--n-text-color);
 background-color: var(--n-color);
 box-sizing: border-box;
 position: relative;
 z-index: auto;
 flex: auto;
 overflow: hidden;
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
`,[(0,h.cB)("layout-scroll-container",`
 overflow-x: hidden;
 box-sizing: border-box;
 height: 100%;
 `),(0,h.cM)("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),m={embedded:Boolean,position:v.u,nativeScrollbar:{type:Boolean,default:!0},scrollbarProps:Object,onScroll:Function,contentClass:String,contentStyle:{type:[String,Object],default:""},hasSider:Boolean,siderPlacement:{type:String,default:"left"}},p=(0,c.U)("n-layout"),g=(0,r.aZ)({name:"Layout",props:Object.assign(Object.assign({},i.Z.props),m),setup(e){let o=(0,r.iH)(null),t=(0,r.iH)(null),{mergedClsPrefixRef:l,inlineThemeDisabled:c}=(0,n.ZP)(e),v=(0,i.Z)("Layout","-layout",u,s.Z,e,l);(0,r.JJ)(p,e);let h=0,m=0;(0,d.M)(()=>{if(e.nativeScrollbar){let e=o.value;e&&(e.scrollTop=m,e.scrollLeft=h)}});let g=(0,r.Fl)(()=>{let{common:{cubicBezierEaseInOut:o},self:t}=v.value;return{"--n-bezier":o,"--n-color":e.embedded?t.colorEmbedded:t.color,"--n-text-color":t.textColor}}),b=c?(0,a.F)("layout",(0,r.Fl)(()=>e.embedded?"e":""),g,e):void 0;return Object.assign({mergedClsPrefix:l,scrollableElRef:o,scrollbarInstRef:t,hasSiderStyle:{display:"flex",flexWrap:"nowrap",width:"100%",flexDirection:"row"},mergedTheme:v,handleNativeElScroll:o=>{var t;let r=o.target;h=r.scrollLeft,m=r.scrollTop,null===(t=e.onScroll)||void 0===t||t.call(e,o)},cssVars:c?void 0:g,themeClass:null==b?void 0:b.themeClass,onRender:null==b?void 0:b.onRender},{scrollTo:function(r,l){if(e.nativeScrollbar){let{value:e}=o;e&&(void 0===l?e.scrollTo(r):e.scrollTo(r,l))}else{let{value:e}=t;e&&e.scrollTo(r,l)}}})},render(){var e;let{mergedClsPrefix:o,hasSider:t}=this;null===(e=this.onRender)||void 0===e||e.call(this);let i=t?this.hasSiderStyle:void 0,n=[this.themeClass,!1,`${o}-layout`,`${o}-layout--${this.position}-positioned`];return(0,r.h)("div",{class:n,style:this.cssVars},this.nativeScrollbar?(0,r.h)("div",{ref:"scrollableElRef",class:[`${o}-layout-scroll-container`,this.contentClass],style:[this.contentStyle,i],onScroll:this.handleNativeElScroll},this.$slots):(0,r.h)(l.Z,Object.assign({},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,contentClass:this.contentClass,contentStyle:[this.contentStyle,i]}),this.$slots))}})},7768:function(e,o,t){t.d(o,{Z:()=>h});var r=t(209),l=t(1321),i=t(4124),n=t(6169),a=t(5291),c=t(2988),d=t(2249);let s=(0,d.cB)("layout-header",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 box-sizing: border-box;
 width: 100%;
 background-color: var(--n-color);
 color: var(--n-text-color);
`,[(0,d.cM)("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 `),(0,d.cM)("bordered",`
 border-bottom: solid 1px var(--n-border-color);
 `)]),v={position:c.u,inverted:Boolean,bordered:{type:Boolean,default:!1}},h=(0,r.aZ)({name:"LayoutHeader",props:Object.assign(Object.assign({},l.Z.props),v),setup(e){let{mergedClsPrefixRef:o,inlineThemeDisabled:t}=(0,i.ZP)(e),c=(0,l.Z)("Layout","-layout-header",s,a.Z,e,o),d=(0,r.Fl)(()=>{let{common:{cubicBezierEaseInOut:o},self:t}=c.value,r={"--n-bezier":o};return e.inverted?(r["--n-color"]=t.headerColorInverted,r["--n-text-color"]=t.textColorInverted,r["--n-border-color"]=t.headerBorderColorInverted):(r["--n-color"]=t.headerColor,r["--n-text-color"]=t.textColor,r["--n-border-color"]=t.headerBorderColor),r}),v=t?(0,n.F)("layout-header",(0,r.Fl)(()=>e.inverted?"a":"b"),d,e):void 0;return{mergedClsPrefix:o,cssVars:t?void 0:d,themeClass:null==v?void 0:v.themeClass,onRender:null==v?void 0:v.onRender}},render(){var e;let{mergedClsPrefix:o}=this;return null===(e=this.onRender)||void 0===e||e.call(this),(0,r.h)("div",{class:[`${o}-layout-header`,this.themeClass,this.position&&`${o}-layout-header--${this.position}-positioned`,this.bordered&&`${o}-layout-header--bordered`],style:this.cssVars},this.$slots)}})},233:function(e,o,t){t.d(o,{Z:()=>z});var r=t(9226),l=t(209),i=t(2121),n=t(1321),a=t(4124),c=t(6169),d=t(3987),s=t(1844),v=t(5568),h=t(5291),u=t(2988),m=t(7610),p=t(2249);let g=(0,p.cB)("layout-sider",`
 flex-shrink: 0;
 box-sizing: border-box;
 position: relative;
 z-index: 1;
 color: var(--n-text-color);
 transition:
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 min-width .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 transform .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 background-color: var(--n-color);
 display: flex;
 justify-content: flex-end;
`,[(0,p.cM)("bordered",[(0,p.cE)("border",`
 content: "";
 position: absolute;
 top: 0;
 bottom: 0;
 width: 1px;
 background-color: var(--n-border-color);
 transition: background-color .3s var(--n-bezier);
 `)]),(0,p.cE)("left-placement",[(0,p.cM)("bordered",[(0,p.cE)("border",`
 right: 0;
 `)])]),(0,p.cM)("right-placement",`
 justify-content: flex-start;
 `,[(0,p.cM)("bordered",[(0,p.cE)("border",`
 left: 0;
 `)]),(0,p.cM)("collapsed",[(0,p.cB)("layout-toggle-button",[(0,p.cB)("base-icon",`
 transform: rotate(180deg);
 `)]),(0,p.cB)("layout-toggle-bar",[(0,p.c)("&:hover",[(0,p.cE)("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),(0,p.cE)("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])])]),(0,p.cB)("layout-toggle-button",`
 left: 0;
 transform: translateX(-50%) translateY(-50%);
 `,[(0,p.cB)("base-icon",`
 transform: rotate(0);
 `)]),(0,p.cB)("layout-toggle-bar",`
 left: -28px;
 transform: rotate(180deg);
 `,[(0,p.c)("&:hover",[(0,p.cE)("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),(0,p.cE)("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})])])]),(0,p.cM)("collapsed",[(0,p.cB)("layout-toggle-bar",[(0,p.c)("&:hover",[(0,p.cE)("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),(0,p.cE)("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])]),(0,p.cB)("layout-toggle-button",[(0,p.cB)("base-icon",`
 transform: rotate(0);
 `)])]),(0,p.cB)("layout-toggle-button",`
 transition:
 color .3s var(--n-bezier),
 right .3s var(--n-bezier),
 left .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 cursor: pointer;
 width: 24px;
 height: 24px;
 position: absolute;
 top: 50%;
 right: 0;
 border-radius: 50%;
 display: flex;
 align-items: center;
 justify-content: center;
 font-size: 18px;
 color: var(--n-toggle-button-icon-color);
 border: var(--n-toggle-button-border);
 background-color: var(--n-toggle-button-color);
 box-shadow: 0 2px 4px 0px rgba(0, 0, 0, .06);
 transform: translateX(50%) translateY(-50%);
 z-index: 1;
 `,[(0,p.cB)("base-icon",`
 transition: transform .3s var(--n-bezier);
 transform: rotate(180deg);
 `)]),(0,p.cB)("layout-toggle-bar",`
 cursor: pointer;
 height: 72px;
 width: 32px;
 position: absolute;
 top: calc(50% - 36px);
 right: -28px;
 `,[(0,p.cE)("top, bottom",`
 position: absolute;
 width: 4px;
 border-radius: 2px;
 height: 38px;
 left: 14px;
 transition: 
 background-color .3s var(--n-bezier),
 transform .3s var(--n-bezier);
 `),(0,p.cE)("bottom",`
 position: absolute;
 top: 34px;
 `),(0,p.c)("&:hover",[(0,p.cE)("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),(0,p.cE)("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})]),(0,p.cE)("top, bottom",{backgroundColor:"var(--n-toggle-bar-color)"}),(0,p.c)("&:hover",[(0,p.cE)("top, bottom",{backgroundColor:"var(--n-toggle-bar-color-hover)"})])]),(0,p.cE)("border",`
 position: absolute;
 top: 0;
 right: 0;
 bottom: 0;
 width: 1px;
 transition: background-color .3s var(--n-bezier);
 `),(0,p.cB)("layout-sider-scroll-container",`
 flex-grow: 1;
 flex-shrink: 0;
 box-sizing: border-box;
 height: 100%;
 opacity: 0;
 transition: opacity .3s var(--n-bezier);
 max-width: 100%;
 `),(0,p.cM)("show-content",[(0,p.cB)("layout-sider-scroll-container",{opacity:1})]),(0,p.cM)("absolute-positioned",`
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 `)]),b=(0,l.aZ)({props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){let{clsPrefix:e}=this;return(0,l.h)("div",{onClick:this.onClick,class:`${e}-layout-toggle-bar`},(0,l.h)("div",{class:`${e}-layout-toggle-bar__top`}),(0,l.h)("div",{class:`${e}-layout-toggle-bar__bottom`}))}});var C=t(8822),x=t(6500);let f=(0,l.aZ)({name:"LayoutToggleButton",props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){let{clsPrefix:e}=this;return(0,l.h)("div",{class:`${e}-layout-toggle-button`,onClick:this.onClick},(0,l.h)(C.Z,{clsPrefix:e},{default:()=>(0,l.h)(x.Z,null)}))}}),y={position:u.u,bordered:Boolean,collapsedWidth:{type:Number,default:48},width:{type:[Number,String],default:272},contentClass:String,contentStyle:{type:[String,Object],default:""},collapseMode:{type:String,default:"transform"},collapsed:{type:Boolean,default:void 0},defaultCollapsed:Boolean,showCollapsedContent:{type:Boolean,default:!0},showTrigger:{type:[Boolean,String],default:!1},nativeScrollbar:{type:Boolean,default:!0},inverted:Boolean,scrollbarProps:Object,triggerClass:String,triggerStyle:[String,Object],collapsedTriggerClass:String,collapsedTriggerStyle:[String,Object],"onUpdate:collapsed":[Function,Array],onUpdateCollapsed:[Function,Array],onAfterEnter:Function,onAfterLeave:Function,onExpand:[Function,Array],onCollapse:[Function,Array],onScroll:Function},z=(0,l.aZ)({name:"LayoutSider",props:Object.assign(Object.assign({},n.Z.props),y),setup(e){let o=(0,l.f3)(m.mj),t=(0,l.iH)(null),i=(0,l.iH)(null),p=(0,l.iH)(e.defaultCollapsed),b=(0,r.Z)((0,l.Vh)(e,"collapsed"),p),C=(0,l.Fl)(()=>(0,d.N)(b.value?e.collapsedWidth:e.width)),x=(0,l.Fl)(()=>"transform"!==e.collapseMode?{}:{minWidth:(0,d.N)(e.width)}),f=(0,l.Fl)(()=>o?o.siderPlacement:"left"),y=0,z=0;(0,v.M)(()=>{if(e.nativeScrollbar){let e=t.value;e&&(e.scrollTop=z,e.scrollLeft=y)}}),(0,l.JJ)(u.T,{collapsedRef:b,collapseModeRef:(0,l.Vh)(e,"collapseMode")});let{mergedClsPrefixRef:I,inlineThemeDisabled:w}=(0,a.ZP)(e),A=(0,n.Z)("Layout","-layout-sider",g,h.Z,e,I),H=(0,l.Fl)(()=>{let{common:{cubicBezierEaseInOut:o},self:t}=A.value,{siderToggleButtonColor:r,siderToggleButtonBorder:l,siderToggleBarColor:i,siderToggleBarColorHover:n}=t,a={"--n-bezier":o,"--n-toggle-button-color":r,"--n-toggle-button-border":l,"--n-toggle-bar-color":i,"--n-toggle-bar-color-hover":n};return e.inverted?(a["--n-color"]=t.siderColorInverted,a["--n-text-color"]=t.textColorInverted,a["--n-border-color"]=t.siderBorderColorInverted,a["--n-toggle-button-icon-color"]=t.siderToggleButtonIconColorInverted,a.__invertScrollbar=t.__invertScrollbar):(a["--n-color"]=t.siderColor,a["--n-text-color"]=t.textColor,a["--n-border-color"]=t.siderBorderColor,a["--n-toggle-button-icon-color"]=t.siderToggleButtonIconColor),a}),S=w?(0,c.F)("layout-sider",(0,l.Fl)(()=>e.inverted?"a":"b"),H,e):void 0;return Object.assign({scrollableElRef:t,scrollbarInstRef:i,mergedClsPrefix:I,mergedTheme:A,styleMaxWidth:C,mergedCollapsed:b,scrollContainerStyle:x,siderPlacement:f,handleNativeElScroll:o=>{var t;let r=o.target;y=r.scrollLeft,z=r.scrollTop,null===(t=e.onScroll)||void 0===t||t.call(e,o)},handleTransitionend:function(o){var t,r;"max-width"===o.propertyName&&(b.value?null===(t=e.onAfterLeave)||void 0===t||t.call(e):null===(r=e.onAfterEnter)||void 0===r||r.call(e))},handleTriggerClick:function(){let{"onUpdate:collapsed":o,onUpdateCollapsed:t,onExpand:r,onCollapse:l}=e,{value:i}=b;t&&(0,s.R)(t,!i),o&&(0,s.R)(o,!i),p.value=!i,i?r&&(0,s.R)(r):l&&(0,s.R)(l)},inlineThemeDisabled:w,cssVars:H,themeClass:null==S?void 0:S.themeClass,onRender:null==S?void 0:S.onRender},{scrollTo:function(o,r){if(e.nativeScrollbar){let{value:e}=t;e&&(void 0===r?e.scrollTo(o):e.scrollTo(o,r))}else{let{value:e}=i;e&&e.scrollTo(o,r)}}})},render(){var e;let{mergedClsPrefix:o,mergedCollapsed:t,showTrigger:r}=this;return null===(e=this.onRender)||void 0===e||e.call(this),(0,l.h)("aside",{class:[`${o}-layout-sider`,this.themeClass,`${o}-layout-sider--${this.position}-positioned`,`${o}-layout-sider--${this.siderPlacement}-placement`,this.bordered&&`${o}-layout-sider--bordered`,t&&`${o}-layout-sider--collapsed`,(!t||this.showCollapsedContent)&&`${o}-layout-sider--show-content`],onTransitionend:this.handleTransitionend,style:[this.inlineThemeDisabled?void 0:this.cssVars,{maxWidth:this.styleMaxWidth,width:(0,d.N)(this.width)}]},this.nativeScrollbar?(0,l.h)("div",{class:[`${o}-layout-sider-scroll-container`,this.contentClass],onScroll:this.handleNativeElScroll,style:[this.scrollContainerStyle,{overflow:"auto"},this.contentStyle],ref:"scrollableElRef"},this.$slots):(0,l.h)(i.Z,Object.assign({},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",style:this.scrollContainerStyle,contentStyle:this.contentStyle,contentClass:this.contentClass,theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,builtinThemeOverrides:this.inverted&&"true"===this.cssVars.__invertScrollbar?{colorHover:"rgba(255, 255, 255, .4)",color:"rgba(255, 255, 255, .3)"}:void 0}),this.$slots),r?"bar"===r?(0,l.h)(b,{clsPrefix:o,class:t?this.collapsedTriggerClass:this.triggerClass,style:t?this.collapsedTriggerStyle:this.triggerStyle,onClick:this.handleTriggerClick}):(0,l.h)(f,{clsPrefix:o,class:t?this.collapsedTriggerClass:this.triggerClass,style:t?this.collapsedTriggerStyle:this.triggerStyle,onClick:this.handleTriggerClick}):null,this.bordered?(0,l.h)("div",{class:`${o}-layout-sider__border`}):null)}})},2988:function(e,o,t){t.d(o,{T:()=>r,u:()=>l});let r=(0,t(1579).U)("n-layout-sider"),l={type:String,default:"static"}},5291:function(e,o,t){t.d(o,{Z:()=>a});var r=t(363),l=t(4738),i=t(1321),n=t(8755);let a=(0,i.j)({name:"Layout",common:n.Z,peers:{Scrollbar:l.Z},self:function(e){let{baseColor:o,textColor2:t,bodyColor:l,cardColor:i,dividerColor:n,actionColor:a,scrollbarColor:c,scrollbarColorHover:d,invertedColor:s}=e;return{textColor:t,textColorInverted:"#FFF",color:l,colorEmbedded:a,headerColor:i,headerColorInverted:s,footerColor:a,footerColorInverted:s,headerBorderColor:n,headerBorderColorInverted:s,footerBorderColor:n,footerBorderColorInverted:s,siderBorderColor:n,siderBorderColorInverted:s,siderColor:i,siderColorInverted:s,siderToggleButtonBorder:`1px solid ${n}`,siderToggleButtonColor:o,siderToggleButtonIconColor:t,siderToggleButtonIconColorInverted:t,siderToggleBarColor:(0,r.h$)(l,c),siderToggleBarColorHover:(0,r.h$)(l,d),__invertScrollbar:"true"}}})},4933:function(e,o,t){t.d(o,{Z:()=>eo});var r=t(7102),l=t(772),i=t(9226),n=t(2370),a=t(209),c=t(8116),d=t(9079),s=t(1321),v=t(4124),h=t(6169),u=t(1844),m=t(2988),p=t(363),g=t(8755),b=t(4440),C=t(6175);let x=(0,s.j)({name:"Menu",common:g.Z,peers:{Tooltip:C.Z,Dropdown:b.Z},self:function(e){var o,t;let{borderRadius:r,textColor3:l,primaryColor:i,textColor2:n,textColor1:a,fontSize:c,dividerColor:d,hoverColor:s,primaryColorHover:v}=e;return Object.assign({borderRadius:r,color:"#0000",groupTextColor:l,itemColorHover:s,itemColorActive:(0,p.zX)(i,{alpha:.1}),itemColorActiveHover:(0,p.zX)(i,{alpha:.1}),itemColorActiveCollapsed:(0,p.zX)(i,{alpha:.1}),itemTextColor:n,itemTextColorHover:n,itemTextColorActive:i,itemTextColorActiveHover:i,itemTextColorChildActive:i,itemTextColorChildActiveHover:i,itemTextColorHorizontal:n,itemTextColorHoverHorizontal:v,itemTextColorActiveHorizontal:i,itemTextColorActiveHoverHorizontal:i,itemTextColorChildActiveHorizontal:i,itemTextColorChildActiveHoverHorizontal:i,itemIconColor:a,itemIconColorHover:a,itemIconColorActive:i,itemIconColorActiveHover:i,itemIconColorChildActive:i,itemIconColorChildActiveHover:i,itemIconColorCollapsed:a,itemIconColorHorizontal:a,itemIconColorHoverHorizontal:v,itemIconColorActiveHorizontal:i,itemIconColorActiveHoverHorizontal:i,itemIconColorChildActiveHorizontal:i,itemIconColorChildActiveHoverHorizontal:i,itemHeight:"42px",arrowColor:n,arrowColorHover:n,arrowColorActive:i,arrowColorActiveHover:i,arrowColorChildActive:i,arrowColorChildActiveHover:i,colorInverted:"#0000",borderColorHorizontal:"#0000",fontSize:c,dividerColor:d},{itemColorHoverInverted:"#0000",itemColorActiveInverted:i,itemColorActiveHoverInverted:i,itemColorActiveCollapsedInverted:i,itemTextColorInverted:o="#BBB",itemTextColorHoverInverted:t="#FFF",itemTextColorChildActiveInverted:t,itemTextColorChildActiveHoverInverted:t,itemTextColorActiveInverted:t,itemTextColorActiveHoverInverted:t,itemTextColorHorizontalInverted:o,itemTextColorHoverHorizontalInverted:t,itemTextColorChildActiveHorizontalInverted:t,itemTextColorChildActiveHoverHorizontalInverted:t,itemTextColorActiveHorizontalInverted:t,itemTextColorActiveHoverHorizontalInverted:t,itemIconColorInverted:o,itemIconColorHoverInverted:t,itemIconColorActiveInverted:t,itemIconColorActiveHoverInverted:t,itemIconColorChildActiveInverted:t,itemIconColorChildActiveHoverInverted:t,itemIconColorCollapsedInverted:o,itemIconColorHorizontalInverted:o,itemIconColorHoverHorizontalInverted:t,itemIconColorActiveHorizontalInverted:t,itemIconColorActiveHoverHorizontalInverted:t,itemIconColorChildActiveHorizontalInverted:t,itemIconColorChildActiveHoverHorizontalInverted:t,arrowColorInverted:o,arrowColorHoverInverted:t,arrowColorActiveInverted:t,arrowColorActiveHoverInverted:t,arrowColorChildActiveInverted:t,arrowColorChildActiveHoverInverted:t,groupTextColorInverted:"#AAA"})}});var f=t(1579);let y=(0,f.U)("n-menu"),z=(0,f.U)("n-submenu"),I=(0,f.U)("n-menu-item-group");var w=t(9065),A=t(2249);let H=[(0,A.c)("&::before","background-color: var(--n-item-color-hover);"),(0,A.cE)("arrow",`
 color: var(--n-arrow-color-hover);
 `),(0,A.cE)("icon",`
 color: var(--n-item-icon-color-hover);
 `),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-hover);
 `,[(0,A.c)("a",`
 color: var(--n-item-text-color-hover);
 `),(0,A.cE)("extra",`
 color: var(--n-item-text-color-hover);
 `)])],S=[(0,A.cE)("icon",`
 color: var(--n-item-icon-color-hover-horizontal);
 `),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-hover-horizontal);
 `,[(0,A.c)("a",`
 color: var(--n-item-text-color-hover-horizontal);
 `),(0,A.cE)("extra",`
 color: var(--n-item-text-color-hover-horizontal);
 `)])],T=(0,A.c)([(0,A.cB)("menu",`
 background-color: var(--n-color);
 color: var(--n-item-text-color);
 overflow: hidden;
 transition: background-color .3s var(--n-bezier);
 box-sizing: border-box;
 font-size: var(--n-font-size);
 padding-bottom: 6px;
 `,[(0,A.cM)("horizontal",`
 max-width: 100%;
 width: 100%;
 display: flex;
 overflow: hidden;
 padding-bottom: 0;
 `,[(0,A.cB)("submenu","margin: 0;"),(0,A.cB)("menu-item","margin: 0;"),(0,A.cB)("menu-item-content",`
 padding: 0 20px;
 border-bottom: 2px solid #0000;
 `,[(0,A.c)("&::before","display: none;"),(0,A.cM)("selected","border-bottom: 2px solid var(--n-border-color-horizontal)")]),(0,A.cB)("menu-item-content",[(0,A.cM)("selected",[(0,A.cE)("icon","color: var(--n-item-icon-color-active-horizontal);"),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-active-horizontal);
 `,[(0,A.c)("a","color: var(--n-item-text-color-active-horizontal);"),(0,A.cE)("extra","color: var(--n-item-text-color-active-horizontal);")])]),(0,A.cM)("child-active",`
 border-bottom: 2px solid var(--n-border-color-horizontal);
 `,[(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-horizontal);
 `,[(0,A.c)("a",`
 color: var(--n-item-text-color-child-active-horizontal);
 `),(0,A.cE)("extra",`
 color: var(--n-item-text-color-child-active-horizontal);
 `)]),(0,A.cE)("icon",`
 color: var(--n-item-icon-color-child-active-horizontal);
 `)]),(0,A.u4)("disabled",[(0,A.u4)("selected, child-active",[(0,A.c)("&:focus-within",S)]),(0,A.cM)("selected",[B(null,[(0,A.cE)("icon","color: var(--n-item-icon-color-active-hover-horizontal);"),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover-horizontal);
 `,[(0,A.c)("a","color: var(--n-item-text-color-active-hover-horizontal);"),(0,A.cE)("extra","color: var(--n-item-text-color-active-hover-horizontal);")])])]),(0,A.cM)("child-active",[B(null,[(0,A.cE)("icon","color: var(--n-item-icon-color-child-active-hover-horizontal);"),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover-horizontal);
 `,[(0,A.c)("a","color: var(--n-item-text-color-child-active-hover-horizontal);"),(0,A.cE)("extra","color: var(--n-item-text-color-child-active-hover-horizontal);")])])]),B("border-bottom: 2px solid var(--n-border-color-horizontal);",S)]),(0,A.cB)("menu-item-content-header",[(0,A.c)("a","color: var(--n-item-text-color-horizontal);")])])]),(0,A.u4)("responsive",[(0,A.cB)("menu-item-content-header",`
 overflow: hidden;
 text-overflow: ellipsis;
 `)]),(0,A.cM)("collapsed",[(0,A.cB)("menu-item-content",[(0,A.cM)("selected",[(0,A.c)("&::before",`
 background-color: var(--n-item-color-active-collapsed) !important;
 `)]),(0,A.cB)("menu-item-content-header","opacity: 0;"),(0,A.cE)("arrow","opacity: 0;"),(0,A.cE)("icon","color: var(--n-item-icon-color-collapsed);")])]),(0,A.cB)("menu-item",`
 height: var(--n-item-height);
 margin-top: 6px;
 position: relative;
 `),(0,A.cB)("menu-item-content",`
 box-sizing: border-box;
 line-height: 1.75;
 height: 100%;
 display: grid;
 grid-template-areas: "icon content arrow";
 grid-template-columns: auto 1fr auto;
 align-items: center;
 cursor: pointer;
 position: relative;
 padding-right: 18px;
 transition:
 background-color .3s var(--n-bezier),
 padding-left .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[(0,A.c)("> *","z-index: 1;"),(0,A.c)("&::before",`
 z-index: auto;
 content: "";
 background-color: #0000;
 position: absolute;
 left: 8px;
 right: 8px;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),(0,A.cM)("disabled",`
 opacity: .45;
 cursor: not-allowed;
 `),(0,A.cM)("collapsed",[(0,A.cE)("arrow","transform: rotate(0);")]),(0,A.cM)("selected",[(0,A.c)("&::before","background-color: var(--n-item-color-active);"),(0,A.cE)("arrow","color: var(--n-arrow-color-active);"),(0,A.cE)("icon","color: var(--n-item-icon-color-active);"),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-active);
 `,[(0,A.c)("a","color: var(--n-item-text-color-active);"),(0,A.cE)("extra","color: var(--n-item-text-color-active);")])]),(0,A.cM)("child-active",[(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-child-active);
 `,[(0,A.c)("a",`
 color: var(--n-item-text-color-child-active);
 `),(0,A.cE)("extra",`
 color: var(--n-item-text-color-child-active);
 `)]),(0,A.cE)("arrow",`
 color: var(--n-arrow-color-child-active);
 `),(0,A.cE)("icon",`
 color: var(--n-item-icon-color-child-active);
 `)]),(0,A.u4)("disabled",[(0,A.u4)("selected, child-active",[(0,A.c)("&:focus-within",H)]),(0,A.cM)("selected",[B(null,[(0,A.cE)("arrow","color: var(--n-arrow-color-active-hover);"),(0,A.cE)("icon","color: var(--n-item-icon-color-active-hover);"),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover);
 `,[(0,A.c)("a","color: var(--n-item-text-color-active-hover);"),(0,A.cE)("extra","color: var(--n-item-text-color-active-hover);")])])]),(0,A.cM)("child-active",[B(null,[(0,A.cE)("arrow","color: var(--n-arrow-color-child-active-hover);"),(0,A.cE)("icon","color: var(--n-item-icon-color-child-active-hover);"),(0,A.cB)("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover);
 `,[(0,A.c)("a","color: var(--n-item-text-color-child-active-hover);"),(0,A.cE)("extra","color: var(--n-item-text-color-child-active-hover);")])])]),(0,A.cM)("selected",[B(null,[(0,A.c)("&::before","background-color: var(--n-item-color-active-hover);")])]),B(null,H)]),(0,A.cE)("icon",`
 grid-area: icon;
 color: var(--n-item-icon-color);
 transition:
 color .3s var(--n-bezier),
 font-size .3s var(--n-bezier),
 margin-right .3s var(--n-bezier);
 box-sizing: content-box;
 display: inline-flex;
 align-items: center;
 justify-content: center;
 `),(0,A.cE)("arrow",`
 grid-area: arrow;
 font-size: 16px;
 color: var(--n-arrow-color);
 transform: rotate(180deg);
 opacity: 1;
 transition:
 color .3s var(--n-bezier),
 transform 0.2s var(--n-bezier),
 opacity 0.2s var(--n-bezier);
 `),(0,A.cB)("menu-item-content-header",`
 grid-area: content;
 transition:
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 opacity: 1;
 white-space: nowrap;
 color: var(--n-item-text-color);
 `,[(0,A.c)("a",`
 outline: none;
 text-decoration: none;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `,[(0,A.c)("&::before",`
 content: "";
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),(0,A.cE)("extra",`
 font-size: .93em;
 color: var(--n-group-text-color);
 transition: color .3s var(--n-bezier);
 `)])]),(0,A.cB)("submenu",`
 cursor: pointer;
 position: relative;
 margin-top: 6px;
 `,[(0,A.cB)("menu-item-content",`
 height: var(--n-item-height);
 `),(0,A.cB)("submenu-children",`
 overflow: hidden;
 padding: 0;
 `,[(0,w.Y)({duration:".2s"})])]),(0,A.cB)("menu-item-group",[(0,A.cB)("menu-item-group-title",`
 margin-top: 6px;
 color: var(--n-group-text-color);
 cursor: default;
 font-size: .93em;
 height: 36px;
 display: flex;
 align-items: center;
 transition:
 padding-left .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `)])]),(0,A.cB)("menu-tooltip",[(0,A.c)("a",`
 color: inherit;
 text-decoration: none;
 `)]),(0,A.cB)("menu-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 6px 18px;
 `)]);function B(e,o){return[(0,A.cM)("hover",e,o),(0,A.c)("&:hover",e,o)]}var E=t(1367),F=t(6985),P=t(1140),N=t(1211),k=t(8822);let R=(0,a.aZ)({name:"ChevronDownFilled",render:()=>(0,a.h)("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},(0,a.h)("path",{d:"M3.20041 5.73966C3.48226 5.43613 3.95681 5.41856 4.26034 5.70041L8 9.22652L11.7397 5.70041C12.0432 5.41856 12.5177 5.43613 12.7996 5.73966C13.0815 6.0432 13.0639 6.51775 12.7603 6.7996L8.51034 10.7996C8.22258 11.0668 7.77743 11.0668 7.48967 10.7996L3.23966 6.7996C2.93613 6.51775 2.91856 6.0432 3.20041 5.73966Z",fill:"currentColor"}))});var $=t(3772);let O=(0,a.aZ)({name:"MenuOptionContent",props:{collapsed:Boolean,disabled:Boolean,title:[String,Function],icon:Function,extra:[String,Function],showArrow:Boolean,childActive:Boolean,hover:Boolean,paddingLeft:Number,selected:Boolean,maxIconSize:{type:Number,required:!0},activeIconSize:{type:Number,required:!0},iconMarginRight:{type:Number,required:!0},clsPrefix:{type:String,required:!0},onClick:Function,tmNode:{type:Object,required:!0},isEllipsisPlaceholder:Boolean},setup(e){let{props:o}=(0,a.f3)(y);return{menuProps:o,style:(0,a.Fl)(()=>{let{paddingLeft:o}=e;return{paddingLeft:o&&`${o}px`}}),iconStyle:(0,a.Fl)(()=>{let{maxIconSize:o,activeIconSize:t,iconMarginRight:r}=e;return{width:`${o}px`,height:`${o}px`,fontSize:`${t}px`,marginRight:`${r}px`}})}},render(){let{clsPrefix:e,tmNode:o,menuProps:{renderIcon:t,renderLabel:r,renderExtra:l,expandIcon:i}}=this,n=t?t(o.rawNode):(0,$.s)(this.icon);return(0,a.h)("div",{onClick:e=>{var o;null===(o=this.onClick)||void 0===o||o.call(this,e)},role:"none",class:[`${e}-menu-item-content`,{[`${e}-menu-item-content--selected`]:this.selected,[`${e}-menu-item-content--collapsed`]:this.collapsed,[`${e}-menu-item-content--child-active`]:this.childActive,[`${e}-menu-item-content--disabled`]:this.disabled,[`${e}-menu-item-content--hover`]:this.hover}],style:this.style},n&&(0,a.h)("div",{class:`${e}-menu-item-content__icon`,style:this.iconStyle,role:"none"},[n]),(0,a.h)("div",{class:`${e}-menu-item-content-header`,role:"none"},this.isEllipsisPlaceholder?this.title:r?r(o.rawNode):(0,$.s)(this.title),this.extra||l?(0,a.h)("span",{class:`${e}-menu-item-content-header__extra`}," ",l?l(o.rawNode):(0,$.s)(this.extra)):null),this.showArrow?(0,a.h)(k.Z,{ariaHidden:!0,class:`${e}-menu-item-content__arrow`,clsPrefix:e},{default:()=>i?i(o.rawNode):(0,a.h)(R,null)}):null)}});function Z(e){let o=(0,a.f3)(y),{props:t,mergedCollapsedRef:r}=o,l=(0,a.f3)(z,null),i=(0,a.f3)(I,null),n=(0,a.Fl)(()=>"horizontal"===t.mode),c=(0,a.Fl)(()=>n.value?t.dropdownPlacement:"tmNodes"in e?"right-start":"right"),d=(0,a.Fl)(()=>{var e;return Math.max(null!==(e=t.collapsedIconSize)&&void 0!==e?e:t.iconSize,t.iconSize)}),s=(0,a.Fl)(()=>{var o;return!n.value&&e.root&&r.value&&null!==(o=t.collapsedIconSize)&&void 0!==o?o:t.iconSize}),v=(0,a.Fl)(()=>{if(n.value)return;let{collapsedWidth:o,indent:a,rootIndent:c}=t,{root:s,isGroup:v}=e;return s?r.value?o/2-d.value/2:void 0===c?a:c:i&&"number"==typeof i.paddingLeftRef.value?a/2+i.paddingLeftRef.value:l&&"number"==typeof l.paddingLeftRef.value?(v?a/2:a)+l.paddingLeftRef.value:0}),h=(0,a.Fl)(()=>{let{collapsedWidth:o,indent:l,rootIndent:i}=t,{value:a}=d,{root:c}=e;return!n.value&&c&&r.value?(void 0===i?l:i)+a+8-(o+a)/2:8});return{dropdownPlacement:c,activeIconSize:s,maxIconSize:d,paddingLeft:v,iconMarginRight:h,NMenu:o,NSubmenu:l}}let M={internalKey:{type:[String,Number],required:!0},root:Boolean,isGroup:Boolean,level:{type:Number,required:!0},title:[String,Function],extra:[String,Function]};var j=t(3898);let L=(0,a.aZ)({name:"MenuDivider",setup(){let{mergedClsPrefixRef:e,isHorizontalRef:o}=(0,a.f3)(y);return()=>o.value?null:(0,a.h)("div",{class:`${e.value}-menu-divider`})}});var _=t(2559);let K=Object.assign(Object.assign({},M),{tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function}),V=(0,P.u)(K),U=(0,a.aZ)({name:"MenuOption",props:K,setup(e){let o=Z(e),{NSubmenu:t,NMenu:r}=o,{props:l,mergedClsPrefixRef:i,mergedCollapsedRef:n}=r,c=t?t.mergedDisabledRef:{value:!1},d=(0,a.Fl)(()=>c.value||e.disabled);return{mergedClsPrefix:i,dropdownPlacement:o.dropdownPlacement,paddingLeft:o.paddingLeft,iconMarginRight:o.iconMarginRight,maxIconSize:o.maxIconSize,activeIconSize:o.activeIconSize,mergedTheme:r.mergedThemeRef,menuProps:l,dropdownEnabled:(0,E.Z)(()=>e.root&&n.value&&"horizontal"!==l.mode&&!d.value),selected:(0,E.Z)(()=>r.mergedValueRef.value===e.internalKey),mergedDisabled:d,handleClick:function(o){d.value||(r.doSelect(e.internalKey,e.tmNode.rawNode),function(o){let{onClick:t}=e;t&&t(o)}(o))}}},render(){let{mergedClsPrefix:e,mergedTheme:o,tmNode:t,menuProps:{renderLabel:r,nodeProps:l}}=this,i=null==l?void 0:l(t.rawNode);return(0,a.h)("div",Object.assign({},i,{role:"menuitem",class:[`${e}-menu-item`,null==i?void 0:i.class]}),(0,a.h)(_.Z,{theme:o.peers.Tooltip,themeOverrides:o.peerOverrides.Tooltip,trigger:"hover",placement:this.dropdownPlacement,disabled:!this.dropdownEnabled||void 0===this.title,internalExtraClass:["menu-tooltip"]},{default:()=>r?r(t.rawNode):(0,$.s)(this.title),trigger:()=>(0,a.h)(O,{tmNode:t,clsPrefix:e,paddingLeft:this.paddingLeft,iconMarginRight:this.iconMarginRight,maxIconSize:this.maxIconSize,activeIconSize:this.activeIconSize,selected:this.selected,title:this.title,extra:this.extra,disabled:this.mergedDisabled,icon:this.icon,onClick:this.handleClick})}))}}),J=Object.assign(Object.assign({},M),{tmNode:{type:Object,required:!0},tmNodes:{type:Array,required:!0}}),D=(0,P.u)(J),q=(0,a.aZ)({name:"MenuOptionGroup",props:J,setup(e){(0,a.JJ)(z,null);let o=Z(e);(0,a.JJ)(I,{paddingLeftRef:o.paddingLeft});let{mergedClsPrefixRef:t,props:r}=(0,a.f3)(y);return function(){let{value:l}=t,i=o.paddingLeft.value,{nodeProps:n}=r,c=null==n?void 0:n(e.tmNode.rawNode);return(0,a.h)("div",{class:`${l}-menu-item-group`,role:"group"},(0,a.h)("div",Object.assign({},c,{class:[`${l}-menu-item-group-title`,null==c?void 0:c.class],style:[(null==c?void 0:c.style)||"",void 0!==i?`padding-left: ${i}px;`:""]}),(0,$.s)(e.title),e.extra?(0,a.h)(a.HY,null," ",(0,$.s)(e.extra)):null),(0,a.h)("div",null,e.tmNodes.map(e=>G(e,r))))}}});function Y(e){return"divider"===e.type||"render"===e.type}function G(e,o){let{rawNode:t}=e,{show:r}=t;if(!1===r)return null;if(Y(t))return"divider"===t.type?(0,a.h)(L,Object.assign({key:e.key},t.props)):null;let{labelField:l}=o,{key:i,level:n,isGroup:c}=e,d=Object.assign(Object.assign({},t),{title:t.title||t[l],extra:t.titleExtra||t.extra,key:i,internalKey:i,level:n,root:0===n,isGroup:c});return e.children?e.isGroup?(0,a.h)(q,(0,j.C)(d,D,{tmNode:e,tmNodes:e.children,key:i})):(0,a.h)(Q,(0,j.C)(d,X,{key:i,rawNodes:t[o.childrenField],tmNodes:e.children,tmNode:e})):(0,a.h)(U,(0,j.C)(d,V,{key:i,tmNode:e}))}let W=Object.assign(Object.assign({},M),{rawNodes:{type:Array,default:()=>[]},tmNodes:{type:Array,default:()=>[]},tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function,domId:String,virtualChildActive:{type:Boolean,default:void 0},isEllipsisPlaceholder:Boolean}),X=(0,P.u)(W),Q=(0,a.aZ)({name:"Submenu",props:W,setup(e){let o=Z(e),{NMenu:t,NSubmenu:r}=o,{props:l,mergedCollapsedRef:i,mergedThemeRef:n}=t,c=(0,a.Fl)(()=>{let{disabled:o}=e;return null!=r&&!!r.mergedDisabledRef.value||!!l.disabled||o}),d=(0,a.iH)(!1);return(0,a.JJ)(z,{paddingLeftRef:o.paddingLeft,mergedDisabledRef:c}),(0,a.JJ)(I,null),{menuProps:l,mergedTheme:n,doSelect:t.doSelect,inverted:t.invertedRef,isHorizontal:t.isHorizontalRef,mergedClsPrefix:t.mergedClsPrefixRef,maxIconSize:o.maxIconSize,activeIconSize:o.activeIconSize,iconMarginRight:o.iconMarginRight,dropdownPlacement:o.dropdownPlacement,dropdownShow:d,paddingLeft:o.paddingLeft,mergedDisabled:c,mergedValue:t.mergedValueRef,childActive:(0,E.Z)(()=>{var o;return null!==(o=e.virtualChildActive)&&void 0!==o?o:t.activePathRef.value.includes(e.internalKey)}),collapsed:(0,a.Fl)(()=>"horizontal"!==l.mode&&(!!i.value||!t.mergedExpandedKeysRef.value.includes(e.internalKey))),dropdownEnabled:(0,a.Fl)(()=>!c.value&&("horizontal"===l.mode||i.value)),handlePopoverShowChange:function(e){d.value=e},handleClick:function(){c.value||(i.value||t.toggleExpand(e.internalKey),function(){let{onClick:o}=e;o&&o()}())}}},render(){var e;let{mergedClsPrefix:o,menuProps:{renderIcon:t,renderLabel:r}}=this,l=()=>{let{isHorizontal:e,paddingLeft:o,collapsed:t,mergedDisabled:r,maxIconSize:l,activeIconSize:i,title:n,childActive:c,icon:d,handleClick:s,menuProps:{nodeProps:v},dropdownShow:h,iconMarginRight:u,tmNode:m,mergedClsPrefix:p,isEllipsisPlaceholder:g,extra:b}=this,C=null==v?void 0:v(m.rawNode);return(0,a.h)("div",Object.assign({},C,{class:[`${p}-menu-item`,null==C?void 0:C.class],role:"menuitem"}),(0,a.h)(O,{tmNode:m,paddingLeft:o,collapsed:t,disabled:r,iconMarginRight:u,maxIconSize:l,activeIconSize:i,title:n,extra:b,showArrow:!e,childActive:c,clsPrefix:p,icon:d,hover:h,onClick:s,isEllipsisPlaceholder:g}))},i=()=>(0,a.h)(F.Z,null,{default:()=>{let{tmNodes:e,collapsed:t}=this;return t?null:(0,a.h)("div",{class:`${o}-submenu-children`,role:"menu"},e.map(e=>G(e,this.menuProps)))}});return this.root?(0,a.h)(N.Z,Object.assign({size:"large",trigger:"hover"},null===(e=this.menuProps)||void 0===e?void 0:e.dropdownProps,{themeOverrides:this.mergedTheme.peerOverrides.Dropdown,theme:this.mergedTheme.peers.Dropdown,builtinThemeOverrides:{fontSizeLarge:"14px",optionIconSizeLarge:"18px"},value:this.mergedValue,disabled:!this.dropdownEnabled,placement:this.dropdownPlacement,keyField:this.menuProps.keyField,labelField:this.menuProps.labelField,childrenField:this.menuProps.childrenField,onUpdateShow:this.handlePopoverShowChange,options:this.rawNodes,onSelect:this.doSelect,inverted:this.inverted,renderIcon:t,renderLabel:r}),{default:()=>(0,a.h)("div",{class:`${o}-submenu`,role:"menu","aria-expanded":!this.collapsed,id:this.domId},l(),this.isHorizontal?null:i())}):(0,a.h)("div",{class:`${o}-submenu`,role:"menu","aria-expanded":!this.collapsed,id:this.domId},l(),i())}}),ee=Object.assign(Object.assign({},s.Z.props),{options:{type:Array,default:()=>[]},collapsed:{type:Boolean,default:void 0},collapsedWidth:{type:Number,default:48},iconSize:{type:Number,default:20},collapsedIconSize:{type:Number,default:24},rootIndent:Number,indent:{type:Number,default:32},labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},disabledField:{type:String,default:"disabled"},defaultExpandAll:Boolean,defaultExpandedKeys:Array,expandedKeys:Array,value:[String,Number],defaultValue:{type:[String,Number],default:null},mode:{type:String,default:"vertical"},watchProps:{type:Array,default:void 0},disabled:Boolean,show:{type:Boolean,default:!0},inverted:Boolean,"onUpdate:expandedKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],onUpdateValue:[Function,Array],"onUpdate:value":[Function,Array],expandIcon:Function,renderIcon:Function,renderLabel:Function,renderExtra:Function,dropdownProps:Object,accordion:Boolean,nodeProps:Function,dropdownPlacement:{type:String,default:"bottom"},responsive:Boolean,items:Array,onOpenNamesChange:[Function,Array],onSelect:[Function,Array],onExpandedNamesChange:[Function,Array],expandedNames:Array,defaultExpandedNames:Array}),eo=(0,a.aZ)({name:"Menu",inheritAttrs:!1,props:ee,setup(e){let{mergedClsPrefixRef:o,inlineThemeDisabled:t}=(0,v.ZP)(e),c=(0,s.Z)("Menu","-menu",T,x,e,o),d=(0,a.f3)(m.T,null),p=(0,a.Fl)(()=>{var o;let{collapsed:t}=e;if(void 0!==t)return t;if(d){let{collapseModeRef:e,collapsedRef:t}=d;if("width"===e.value)return null!==(o=t.value)&&void 0!==o&&o}return!1}),g=(0,a.Fl)(()=>{let{keyField:o,childrenField:t,disabledField:r}=e;return(0,l.J)(e.items||e.options,{getIgnored:e=>Y(e),getChildren:e=>e[t],getDisabled:e=>e[r],getKey(e){var t;return null!==(t=e[o])&&void 0!==t?t:e.name}})}),b=(0,a.Fl)(()=>new Set(g.value.treeNodes.map(e=>e.key))),{watchProps:C}=e,f=(0,a.iH)(null);(null==C?void 0:C.includes("defaultValue"))?(0,a.m0)(()=>{f.value=e.defaultValue}):f.value=e.defaultValue;let z=(0,a.Vh)(e,"value"),I=(0,i.Z)(z,f),w=(0,a.iH)([]),A=()=>{w.value=e.defaultExpandAll?g.value.getNonLeafKeys():e.defaultExpandedNames||e.defaultExpandedKeys||g.value.getPath(I.value,{includeSelf:!1}).keyPath};(null==C?void 0:C.includes("defaultExpandedKeys"))?(0,a.m0)(A):A();let H=(0,n.Z)(e,["expandedNames","expandedKeys"]),S=(0,i.Z)(H,w),B=(0,a.Fl)(()=>g.value.treeNodes),E=(0,a.Fl)(()=>g.value.getPath(I.value).keyPath);function F(o){let{"onUpdate:expandedKeys":t,onUpdateExpandedKeys:r,onExpandedNamesChange:l,onOpenNamesChange:i}=e;t&&(0,u.R)(t,o),r&&(0,u.R)(r,o),l&&(0,u.R)(l,o),i&&(0,u.R)(i,o),w.value=o}(0,a.JJ)(y,{props:e,mergedCollapsedRef:p,mergedThemeRef:c,mergedValueRef:I,mergedExpandedKeysRef:S,activePathRef:E,mergedClsPrefixRef:o,isHorizontalRef:(0,a.Fl)(()=>"horizontal"===e.mode),invertedRef:(0,a.Vh)(e,"inverted"),doSelect:function(o,t){let{"onUpdate:value":r,onUpdateValue:l,onSelect:i}=e;l&&(0,u.R)(l,o,t),r&&(0,u.R)(r,o,t),i&&(0,u.R)(i,o,t),f.value=o},toggleExpand:function(o){let t=Array.from(S.value),r=t.findIndex(e=>e===o);if(~r)t.splice(r,1);else{if(e.accordion&&b.value.has(o)){let e=t.findIndex(e=>b.value.has(e));e>-1&&t.splice(e,1)}t.push(o)}F(t)}});let P=(0,a.Fl)(()=>{let{inverted:o}=e,{common:{cubicBezierEaseInOut:t},self:r}=c.value,{borderRadius:l,borderColorHorizontal:i,fontSize:n,itemHeight:a,dividerColor:d}=r,s={"--n-divider-color":d,"--n-bezier":t,"--n-font-size":n,"--n-border-color-horizontal":i,"--n-border-radius":l,"--n-item-height":a};return o?(s["--n-group-text-color"]=r.groupTextColorInverted,s["--n-color"]=r.colorInverted,s["--n-item-text-color"]=r.itemTextColorInverted,s["--n-item-text-color-hover"]=r.itemTextColorHoverInverted,s["--n-item-text-color-active"]=r.itemTextColorActiveInverted,s["--n-item-text-color-child-active"]=r.itemTextColorChildActiveInverted,s["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveInverted,s["--n-item-text-color-active-hover"]=r.itemTextColorActiveHoverInverted,s["--n-item-icon-color"]=r.itemIconColorInverted,s["--n-item-icon-color-hover"]=r.itemIconColorHoverInverted,s["--n-item-icon-color-active"]=r.itemIconColorActiveInverted,s["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHoverInverted,s["--n-item-icon-color-child-active"]=r.itemIconColorChildActiveInverted,s["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHoverInverted,s["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsedInverted,s["--n-item-text-color-horizontal"]=r.itemTextColorHorizontalInverted,s["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontalInverted,s["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontalInverted,s["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontalInverted,s["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontalInverted,s["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontalInverted,s["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontalInverted,s["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontalInverted,s["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontalInverted,s["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontalInverted,s["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontalInverted,s["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontalInverted,s["--n-arrow-color"]=r.arrowColorInverted,s["--n-arrow-color-hover"]=r.arrowColorHoverInverted,s["--n-arrow-color-active"]=r.arrowColorActiveInverted,s["--n-arrow-color-active-hover"]=r.arrowColorActiveHoverInverted,s["--n-arrow-color-child-active"]=r.arrowColorChildActiveInverted,s["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHoverInverted,s["--n-item-color-hover"]=r.itemColorHoverInverted,s["--n-item-color-active"]=r.itemColorActiveInverted,s["--n-item-color-active-hover"]=r.itemColorActiveHoverInverted,s["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsedInverted):(s["--n-group-text-color"]=r.groupTextColor,s["--n-color"]=r.color,s["--n-item-text-color"]=r.itemTextColor,s["--n-item-text-color-hover"]=r.itemTextColorHover,s["--n-item-text-color-active"]=r.itemTextColorActive,s["--n-item-text-color-child-active"]=r.itemTextColorChildActive,s["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveHover,s["--n-item-text-color-active-hover"]=r.itemTextColorActiveHover,s["--n-item-icon-color"]=r.itemIconColor,s["--n-item-icon-color-hover"]=r.itemIconColorHover,s["--n-item-icon-color-active"]=r.itemIconColorActive,s["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHover,s["--n-item-icon-color-child-active"]=r.itemIconColorChildActive,s["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHover,s["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsed,s["--n-item-text-color-horizontal"]=r.itemTextColorHorizontal,s["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontal,s["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontal,s["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontal,s["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontal,s["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontal,s["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontal,s["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontal,s["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontal,s["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontal,s["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontal,s["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontal,s["--n-arrow-color"]=r.arrowColor,s["--n-arrow-color-hover"]=r.arrowColorHover,s["--n-arrow-color-active"]=r.arrowColorActive,s["--n-arrow-color-active-hover"]=r.arrowColorActiveHover,s["--n-arrow-color-child-active"]=r.arrowColorChildActive,s["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHover,s["--n-item-color-hover"]=r.itemColorHover,s["--n-item-color-active"]=r.itemColorActive,s["--n-item-color-active-hover"]=r.itemColorActiveHover,s["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsed),s}),N=t?(0,h.F)("menu",(0,a.Fl)(()=>e.inverted?"a":"b"),P,e):void 0,k=(0,r.Mc)(),R=(0,a.iH)(null),$=(0,a.iH)(null),O=!0,Z=()=>{var e;O?O=!1:null===(e=R.value)||void 0===e||e.sync({showAllItemsBeforeCalculate:!0})},M=(0,a.iH)(-1),j=(0,a.Fl)(()=>{let o=M.value;return{children:-1===o?[]:e.options.slice(o)}}),L=(0,a.Fl)(()=>{let{childrenField:o,disabledField:t,keyField:r}=e;return(0,l.J)([j.value],{getIgnored:e=>Y(e),getChildren:e=>e[o],getDisabled:e=>e[t],getKey(e){var o;return null!==(o=e[r])&&void 0!==o?o:e.name}})}),_=(0,a.Fl)(()=>(0,l.J)([{}]).treeNodes[0]);return{mergedClsPrefix:o,controlledExpandedKeys:H,uncontrolledExpanededKeys:w,mergedExpandedKeys:S,uncontrolledValue:f,mergedValue:I,activePath:E,tmNodes:B,mergedTheme:c,mergedCollapsed:p,cssVars:t?void 0:P,themeClass:null==N?void 0:N.themeClass,overflowRef:R,counterRef:$,updateCounter:()=>{},onResize:Z,onUpdateOverflow:function(e){e||(M.value=-1)},onUpdateCount:function(o){M.value=e.options.length-o},renderCounter:function(){var e;if(-1===M.value)return(0,a.h)(Q,{root:!0,level:0,key:"__ellpisisGroupPlaceholder__",internalKey:"__ellpisisGroupPlaceholder__",title:"\xb7\xb7\xb7",tmNode:_.value,domId:k,isEllipsisPlaceholder:!0});let o=L.value.treeNodes[0],t=E.value,r=!!(null===(e=o.children)||void 0===e?void 0:e.some(e=>t.includes(e.key)));return(0,a.h)(Q,{level:0,root:!0,key:"__ellpisisGroup__",internalKey:"__ellpisisGroup__",title:"\xb7\xb7\xb7",virtualChildActive:r,tmNode:o,domId:k,rawNodes:o.rawNode.children||[],tmNodes:o.children||[],isEllipsisPlaceholder:!0})},getCounter:function(){return document.getElementById(k)},onRender:null==N?void 0:N.onRender,showOption:o=>{let t=g.value.getPath(null!=o?o:I.value,{includeSelf:!1}).keyPath;if(!t.length)return;let r=new Set([...Array.from(S.value),...t]);e.accordion&&b.value.forEach(e=>{r.has(e)&&!t.includes(e)&&r.delete(e)}),F(Array.from(r))},deriveResponsiveState:Z}},render(){let{mergedClsPrefix:e,mode:o,themeClass:t,onRender:r}=this;null==r||r();let l=()=>this.tmNodes.map(e=>G(e,this.$props)),i="horizontal"===o&&this.responsive,n=()=>(0,a.h)("div",(0,a.dG)(this.$attrs,{role:"horizontal"===o?"menubar":"menu",class:[`${e}-menu`,t,`${e}-menu--${o}`,i&&`${e}-menu--responsive`,this.mergedCollapsed&&`${e}-menu--collapsed`],style:this.cssVars}),i?(0,a.h)(c.Z,{ref:"overflowRef",onUpdateOverflow:this.onUpdateOverflow,getCounter:this.getCounter,onUpdateCount:this.onUpdateCount,updateCounter:this.updateCounter,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:l,counter:this.renderCounter}):l());return i?(0,a.h)(d.Z,{onResize:this.onResize},{default:n}):n()}})}}]);