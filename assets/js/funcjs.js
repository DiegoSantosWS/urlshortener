$(document).ready(function(){
    //alert("teste")
    showresults();
});

function showresults() {
    $("#showresults").hide(100)
    //Carrega Lista de conteudos cadastrados
    $.ajax({
        url: "/list",
        type:"jsonp",
        crossDomain: true,
        success:function(data) {
            var html = "";
            var str = "";
            var urlNew = "";
            setTimeout(function() {
                jQuery.each(data, function(i, item){
                    str = item.url;
                    if (str.length >= 100) {
                        urlNew = str.substr(0,100) +"...";
                    } else {
                        urlNew = item.url
                    }
                    html += "<tr>";
                    html += "<td><a href='"+item.url+"'>"+urlNew+"</a></td>";
                    html += "<td>";
                    html += "<a id='"+i+"' href='http://localhost:3000/r/"+item.token+"'>http://localhost:3000/r/"+item.token+"</a>";
                    html += "   <button onclick='copyToClipboard("+i+")' title='Copy short URL'><i class='fa fa-copy fa-2 text-primary' aria-hidden='true'></i></button>";
                    html += "</td>";
                    html += "<td>"+item.total+"</td>";
                    html += "<td>";
                    html += '<a class="btn btn-warning" href="http://localhost:3000/info/'+item.token+'" title="Views Information of access url"><i class="fa fa-eye fa-2 text-primary" aria-hidden="true"></i></a>';
                    html += "</td>";
                    html += "</tr>"; 
                })
                $("#showresults").show(100)
                $("#viwsResults").html(html)    
            },10);
        }
    });
}

function viewIformation(token) {
    alert(token)
}

function copyToClipboard(elementId) {
    // Create a "hidden" input
    var aux = document.createElement("input");
    // Assign it the value of the specified element
    aux.setAttribute("value", document.getElementById(elementId).innerHTML);
    // Append it to the body
    document.body.appendChild(aux);
    // Highlight its content
    aux.select();
  
    // Copy the highlighted text
    document.execCommand("copy");
  
    // Remove it from the body
    document.body.removeChild(aux);
}