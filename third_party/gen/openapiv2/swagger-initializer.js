window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [{"url":"qclaogui/routeguide/v1/route_guide.swagger.json","name":"qclaogui/routeguide/v1/route_guide.swagger.json"},{"url":"qclaogui/bookstore/v1alpha1/bookstore.swagger.json","name":"qclaogui/bookstore/v1alpha1/bookstore.swagger.json"},{"url":"qclaogui/library/v1/service.swagger.json","name":"qclaogui/library/v1/service.swagger.json"},{"url":"qclaogui/library/v1/shelf.swagger.json","name":"qclaogui/library/v1/shelf.swagger.json"},{"url":"qclaogui/library/v1/book.swagger.json","name":"qclaogui/library/v1/book.swagger.json"},{"url":"qclaogui/todo/v1/todo_service.swagger.json","name":"qclaogui/todo/v1/todo_service.swagger.json"}],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
