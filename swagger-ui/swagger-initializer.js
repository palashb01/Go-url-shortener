window.onload = function () {
    window.ui = SwaggerUIBundle({
      url: "/docs/shortener.swagger.json",
      dom_id: "#swagger-ui",
      deepLinking: true,
      presets: [SwaggerUIBundle.presets.apis],
      validatorUrl: null,
      // Point Swagger to the **actual Go API** (on port 8080)
      requestInterceptor: (req) => {
        req.url = req.url.replace('localhost:8081', 'localhost:8080');
        return req;
      }
    });
  };
  