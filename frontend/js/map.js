function initMap() {
    var opts = {
        zoom: 12,
        //マップの中心座標
        center: new google.maps.LatLng(33.9515, 131.2467)
    };
    var map = new google.maps.Map(document.getElementById("map"), opts);

    //ピンの座標指定
    var m_latlng1 = new google.maps.LatLng(33.9479435,131.2828868);
    var marker1 = new google.maps.Marker({
        position: m_latlng1,
        map: map
    });
    //吹き出し
    var myInfoWindow = new google.maps.InfoWindow({
        // 情報ウインドウの内容かくところ
        content: `<div class="name">
                    田植えイベント
                    </div>
                    <div class="address">
                    常盤公園
                    </div>`
    });
    myInfoWindow.open(map, marker1);
    google.maps.event.addListener(myInfoWindow, "closeclick", function() {
        google.maps.event.addListenerOnce(marker1, "click", function(event) {
            myInfoWindow.open(map, marker1);
        });
    });

    var m_latlng2 = new google.maps.LatLng(33.998403,131.211729);
    var marker2 = new google.maps.Marker({
        position: m_latlng2,
        map: map
    });
    //吹き出し
    var myInfoWindow = new google.maps.InfoWindow({
        // 情報ウインドウの内容かくところ
        content: `<div class="name">
                    もちまきやります！
                    </div>
                    <div class="address">
                    自宅付近
                　　</div>`
    });
    myInfoWindow.open(map, marker2);
    google.maps.event.addListener(myInfoWindow, "closeclick", function() {
        google.maps.event.addListenerOnce(marker2, "click", function(event) {
            myInfoWindow.open(map, marker2);
        });
    });
}