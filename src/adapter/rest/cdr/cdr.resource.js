/**
 * Created by Igor Navrotskyj on 25.08.2015.
 */

'use strict';

var conf = require(__appRoot + '/conf'),
    http = require('http'),
    https = require('https'),
    url = require('url'),
    log = require(__appRoot +  '/lib/log')(module),
    CDR_SERVER_HOST = conf.get('cdrServer:host'),
    middlewareRequest = conf.get('cdrServer:useProxy') ? proxyToCdr : redirectToCdr;
    ;

module.exports = {
    addRoutes: addRoutes
};

/**
 * Adds routes to the api.
 */

function addRoutes(api) {
    api.all(/^\/api\/v2\/r\/(cdr|files|media)/, getRedirectUrl);
    api.all(/^\/api\/v2\/(cdr|files|media)/, middlewareRequest);
};

if (CDR_SERVER_HOST) {
    CDR_SERVER_HOST = CDR_SERVER_HOST.replace(/\/$/g, '');
};

var cdrHostInfo = url.parse(CDR_SERVER_HOST);

var client = cdrHostInfo.protocol == 'http:' ? http.request : https.request;

var CDR_SERVER = {
    path: cdrHostInfo.path.length == 1 ? '' : cdrHostInfo.path,
    hostName: cdrHostInfo.hostname,
    port: parseInt(cdrHostInfo.port)
};

function redirectToCdr(request, response, next) {
    response.redirect(307, CDR_SERVER_HOST + request.originalUrl);
};

function proxyToCdr(request, response, next) {
    var postData = JSON.stringify(request.body);
    var options = {
        hostname: CDR_SERVER.hostName,
        port: CDR_SERVER.port,
        path: CDR_SERVER.path + request.originalUrl,
        headers: {
        },
        method: request.method,
        rejectUnauthorized: false
    };

    for (let key in request.headers)
        options.headers[key] = request.headers[key]

    if (request.headers.hasOwnProperty('content-length')) {
        options.headers['content-length'] = request._body
            ? Buffer.byteLength(postData)
            : request.headers['content-length'];
    };

    var req = client(options, function(res) {
        try {
            res.on('end', function () {
                res.destroy();
            });
            response.writeHead(res.statusCode, res.headers);

            res.pipe(response);
        } catch (e){
            log.error(e);
        }
    });

    req.on('error', function(e) {
        log.error(e);
        next(e);
    });

// write data to request body
    if (request._body) {
        req.write(postData);
    };
    request.on('end', function () {
        req.end();
    });
    request.pipe(req);
};

function getRedirectUrl(req, res, next) {
    if (!CDR_SERVER_HOST) {
        return res.status(500).json({
            "status": "error",
            "info": "Not config CDR_SERVER_HOST"
        });
    };
    res.status(200).json({
        "status": "OK",
        "info": CDR_SERVER_HOST + req.originalUrl.replace(/(\/api\/v2\/)(r\/)/, '$1')
    });
};

function getPort(req) {
    var matches = req.headers.host.match(/:(\d+)/);
    if (matches)
        return matches[1];
    return req.connection.pair ? '443' : '80';
};