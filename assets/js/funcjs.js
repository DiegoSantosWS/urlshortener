$(document).ready(function(){
    
    $("#email").on("change", function(){
        var email = $(this).val();

        $.ajax({
            type: "POST",
            url:"/check-cad/"+email,
            dataType: "json",
            success: function(data) {
                if (data != 0) {
                    swal('OPS:',"O e-mail, "+email+" já existe em nosso sistema",'warning');
                    //alert("O email: "+email+" já existe em nosso sistema");
                    $("#email").val('')
                    $("#email").focus()
                }
            }
        })
    });
    
    showresults();
    viewIformation($("#tokenAnalytcis").val());
    //showChartBrowser($("#tokenAnalytcisChartBrowser").val());
});
/**
 * Abre a aprte de informações de acessos.
 * @param {*} id 
 */
function analytics(id) {
    setTimeout(function(){
        window.location.href = "/analytics-wd/"+id;
    }, 100);
}
/** 
 * Exibe os resultados de cadastros
 */
function showresults() {
    //Carrega Lista de conteudos cadastrados
    $.ajax({
        url: "/list/",
        type:"json",
        crossDomain: true,
        success:function(data) {
            var html = "";
            var str = "";
            var urlNew = "";
            
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
                html += "<a id='"+i+"' href='http://"+item.shortenURL+"'>http://"+item.shortenURL+"</a>";
                html += "   <button onclick='copyToClipboard("+i+")' title='Copy short URL'><i class='fa fa-copy fa-2 text-primary' aria-hidden='true'></i></button>";
                html += "</td>";
                html += "<td>"+item.total+"</td>";
                html += "<td>";
                html += "<button type='button' onclick=\'analytics(\""+item.token+"\")\' class='btn btn-warning' title='Views Information of access url'>";
                html += "<i class='fa fa-eye fa-2 text-primary' aria-hidden='true'></i>";
                html += "</button>";
                html += "</td>";
                html += "</tr>"; 
            });
            $("#viwsResults").html(html);
        }
    });
}
/**
 * Retorna informações de uma url
 * @param {*} cod 
 */
function viewIformation(cod) {
    $.ajax({
        url: "/info/"+cod,
        type:"jsonp",
        crossDomain: true,
        success:function(data) {
            var html = "";
            jQuery.each(data, function(i, item){
                
                html += "<tr>";
                html += "<td>"+item.referencia.String+"</td>";
                html += "<td>"+item.contador+"</td>";
                html += "<td>"+item.browser.String+"</td>";
                html += "<td>"+item.sysoperacional.String+"</td>";
                html += "<td>"+item.data+"</td>";
                html += "</tr>"; 
            })
            $("#viwsResultsAnalytics").html(html);
        }
    });
}
/**
 * Recebe um elemento para ser copiado.
 * @param {*} elementId 
 */
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