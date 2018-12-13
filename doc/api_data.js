define({ "api": [
  {
    "type": "post",
    "url": "/upload/url/:url",
    "title": "Upload Stream File",
    "name": "Upload",
    "group": "upload",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "url",
            "optional": false,
            "field": "url",
            "description": "<p>base64 encoded url.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>return code of the status.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>return message of the status.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "service/route.go",
    "groupTitle": "upload"
  }
] });
