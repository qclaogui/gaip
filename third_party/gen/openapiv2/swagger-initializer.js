window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [{"url":"qclaogui/todo/v1/todo.swagger.json","name":"qclaogui/todo/v1/todo.swagger.json"},{"url":"qclaogui/todo/v1/todo_service.swagger.json","name":"qclaogui/todo/v1/todo_service.swagger.json"}],
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
