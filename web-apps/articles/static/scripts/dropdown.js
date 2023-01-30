$("#ddForm button").click(function(event){
    event.preventDefault();
    var url = $("#ddForm select").val() || false;
    var action = $(this).data("action");
    if(action === "down" && url){
        window.location.href = url;
    } else if(action === "up" && url){
        window.location.href = "http://www.link.com/download.php?url=" + url;
    }
});