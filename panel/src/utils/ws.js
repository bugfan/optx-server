export var ws = {}

export default {
    ws: {},
    initWebSocket: function (){ //初始化weosocket
        var login = this.login
        this.socket = new WebSocket('ws://192.168.31.44/reunion/websocket')
        this.ws = Stomp.over(this.socket)
        var _this=this
        this.ws.connect({
            username: login,
            sessionId: this.$store.state.sessionId
        },
        function(succ) {
            console.log('连接成功:', succ);
            _this.ws.subscribe("/queue/errors", function(message) {
            alert("Error " + message.body) //链接失败
        });
        _this.ws.subscribe("/user/notice/message", function(message) { //接受消息时
            console.log('个人消息：',message)
            _this.showMessage(message.body,"queue") //拿到的消息
        });
        _this.ws.subscribe("/topic/reply/" + _this.to, function (message) {
            console.log('群组消息：',message)
            _this.showMessage(message.body,"topic");
        });
        // alert("登录成功")
        },
        function(error) {
            // alert("创建对象失败 " + error)
        })
    }
}