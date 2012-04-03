//Requires JQuery

//name: name of function to call
//data: json data
//returntype: what is to be expected
//success: function(json) {}
function Communicate(name, data, returntype, success) {
	if (typeof(data) != "string") {
		data = JSON.stringify(data);
	}
	$.ajax({
		url: "/?ajax="+name,
		success: success,
		error: function(a, b, c) {alert("Cannot Connect to server?\n\nor\n\n"+String(a)+String(b), String(c))},
		data: {"data": data},
		dataType: returntype,
		type:"POST"
	});
}

//Server closes after 15seconds without a ping
$.get("/ping");
window.setInterval(function(){
	$.get("/ping");
}, 10*1000)
