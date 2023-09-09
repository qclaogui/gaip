window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [{"url":"routeguide/v1/route_guide.swagger.json","name":"routeguide/v1/route_guide.swagger.json"},{"url":"bookstore/v1beta1/bookstore.swagger.json","name":"bookstore/v1beta1/bookstore.swagger.json"},{"url":"bookstore/v1alpha1/bookstore.swagger.json","name":"bookstore/v1alpha1/bookstore.swagger.json"},{"url":"todo/v1/todo_service.swagger.json","name":"todo/v1/todo_service.swagger.json"}],
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
