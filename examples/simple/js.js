function CLICK() {
	Communicate("test", {"Data":"it works"}, "json", function(data) {
		$("#kage").html(data["Data"]);
	});
}
function Reverse() {
	Communicate("reverse", $("#myinput").val(), "text", function(data) {
		$("#thing").html(data);
	});
} 
