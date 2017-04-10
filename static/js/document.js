layui.define(['layer', 'form', 'element','tree'], function(exports) {
    var layer = layui.layer;
    var elements = layui.element();
    var $ = layui.jquery;
    var form = layui.form();
    layui.tree({
  elem: '#demo' //传入元素选择器
  ,nodes: [{ //节点
    name: '父节点1'
    ,children: [{
      name: '子节点11'
    },{
      name: '子节点12'
    }]
  },{
    name: '父节点2（可以点左侧箭头，也可以双击标题）'
    ,children: [{
      name: '子节点21'
      ,children: [{
        name: '子节点211'
      }]
    }]
  }]
});
    exports('document', {});
});