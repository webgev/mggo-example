var api = function (method, params) {
    params = params || {}
    var settings = {
        "method": "POST",
        "type": "POST",
        "headers": {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        "body": JSON.stringify({
            "method": method,
            "params": params
        })
    };

    return fetch('/api/', settings)
                .then(response=> {
                   
                    return response.json()
                })
                .then(res=>{
                    if (res.Error) {
                        throw new Error(res.Error);
                    }
                    return res.Result
                })
}
var apiFile = function (method, params, image) {
    params = params || {}
    var formData  = new FormData();
    formData.append("method", method)
    formData.append("params", JSON.stringify(params))
    formData.append('file', image)

    var settings = {
        "method": "POST",
        "type": "POST",
        "headers": {
            'accept': "application/json, text/javascript",
        //    'content-type': 'multipart/form-data'
        },
        "body": formData
    };

    return fetch('/api/', settings).then(res=> res.json()).then(res=>res.Result)
}