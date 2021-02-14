var CACHE_NAME = "Spago-sample-v1";

console.log("loading: serviceworker");

var urlsToCache = [
    ".",
    "./wasm_exec.js",
    "./main.wasm",
    "./favicon.ico",
    "./serviceworker.js",
    "./assets/app.css",
    "./assets/threejs/build/three.module.js",
    "https://cdn.jsdelivr.net/npm/uikit@3.6.5/dist/css/uikit.min.css",
    "https://cdn.jsdelivr.net/npm/uikit@3.6.5/dist/js/uikit.min.js",
    "https://cdn.jsdelivr.net/npm/uikit@3.6.5/dist/js/uikit-icons.min.js",
    "./assets/threejs/ex/jsm/controls/OrbitControls.js",
    "./assets/threejs/ex/jsm/effects/OutlineEffect.js",
    "./assets/threejs/ex/jsm/animation/MMDAnimationHelper.js",
    "./assets/threejs/ex/jsm/animation/MMDPhysics.js",
    "./assets/threejs/ex/jsm/animation/CCDIKSolver.js",
    "./assets/threejs/ex/jsm/loaders/MMDLoader.js",
    "./assets/threejs/ex/jsm/libs/mmdparser.module.js",
    "./assets/threejs/ex/jsm/loaders/TGALoader.js",
    "./assets/threejs/ex/jsm/libs/dat.gui.module.js",
    "./assets/threejs/ex/jsm/libs/stats.module.js",
    "./assets/threejs/ex/js/libs/ammo.wasm.js",
    "./assets/threejs/ex/js/libs/ammo.wasm.wasm"
    // "./assets/models/mmd/miku/miku_v2.pmd",
    // "./assets/models/mmd/miku/eyeM2.bmp",
    // "./assets/models/mmd/vmds/wavefile_v2.vmd"
];

// 残したいキャッシュのバージョンをこの配列に入れる
// 基本的に現行の1つだけでよい。他は削除される。
const CACHE_KEYS = [CACHE_NAME];

self.addEventListener("install", function (event) {
    console.log("install: serviceworker");
    event.waitUntil(
        caches
            .open(CACHE_NAME) // 上記で指定しているキャッシュ名
            .then(function (cache) {
                return cache.addAll(urlsToCache); // 指定したリソースをキャッシュへ追加
                // 1つでも失敗したらService Workerのインストールはスキップされる
            })
    );
});

//新しいバージョンのServiceWorkerが有効化されたとき
self.addEventListener("activate", (event) => {
    console.log("activate: serviceworker");
    event.waitUntil(
        caches.keys().then((keys) => {
            return Promise.all(
                keys
                    .filter((key) => {
                        return !CACHE_KEYS.includes(key);
                    })
                    .map((key) => {
                        // 不要なキャッシュを削除
                        return caches.delete(key);
                    })
            );
        })
    );
});

self.addEventListener("fetch", function (event) {
    var online = navigator.onLine;

    // ファイルパス ~/test.htmlにアクセスすると、このファイル自体は無いがServiceWorkerがResponseを作成して表示してくれる
    if (event.request.url.indexOf("test.html") != -1) {
        return event.respondWith(
            new Response("任意のURLの内容をここで自由に返却できる")
        );
    }

    if (online) {
        console.log("ONLINE");
        //このパターンの処理では、Responseだけクローンすれば問題ない
        //cloneEventRequest = event.request.clone();
        event.respondWith(
            caches.match(event.request).then(function (response) {
                if (response) {
                    //オンラインでもローカルにキャッシュでリソースがあればそれを返す
                    //ここを無効にすればオンラインのときは常にオンラインリソースを取りに行き、その最新版をキャッシュにPUTする
                    return response;
                }
                //request streem 1
                return fetch(event.request)
                    .then(function (response) {
                        //ローカルキャッシュになかったからネットワークから落とす
                        //ネットワークから落とせてればここでリソースが返される

                        // Responseはストリームなのでキャッシュで使用してしまうと、ブラウザの表示で不具合が起こる(っぽい)ので、複製したものを使う
                        cloneResponse = response.clone();

                        if (response) {
                            if (response || response.status == 200) {
                                console.log("正常にリソースを取得");
                                caches.open(CACHE_NAME).then(function (cache) {
                                    console.log("キャッシュへ保存");
                                    //初回表示でエラー起きているが致命的でないので保留
                                    cache.put(event.request, cloneResponse).then(function () {
                                        console.log("保存完了");
                                    });
                                });
                            } else {
                                return event.respondWith(
                                    new Response("200以外のエラーをハンドリングしたりできる")
                                );
                            }
                            return response;
                        }
                    })
                    .catch(function (error) {
                        return console.log(error);
                    });
            })
        );
    } else {
        console.log("OFFLINE");
        event.respondWith(
            caches.match(event.request).then(function (response) {
                // キャッシュがあったのでそのレスポンスを返す
                if (response) {
                    return response;
                }
                //オフラインでキャッシュもなかったパターン
                return caches.match("offline.html").then(function (responseNodata) {
                    return responseNodata;
                });
            })
        );
    }
});