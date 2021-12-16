String.prototype.format = function () {
    var args = [].slice.call(arguments);
    return this.replace(/(\{\d+\})/g, function (a){
        return args[+(a.substr(1,a.length-2))||0];
    });
};
const foo = document.getElementById("exp").innerText;
const bar = foo.split(".");
document.getElementById("exp").innerText = "";
for (let i = 0; i < bar.length; i++) {
    document.getElementById("exp").insertAdjacentHTML('beforeend',"<p>" + bar[i] + "</p>");
}
document.getElementById("forward").addEventListener("click", function() {
    navigate("forward");
});
document.getElementById("back").addEventListener("click", function() {
    navigate("back")
});
function navigate(direction) {
    let foo = window.location.href.split("?");
    let quuz = foo[1].split("=");
    let date = new Date(quuz[1]);
    if (direction === "back") {
        date.setDate(date.getDate() - 1);
    } else {
        date.setDate(date.getDate() + 1);
    }
    let next = "{0}-{1}-{2}".format(date.getFullYear(), date.getUTCMonth() + 1, date.getDate());
    window.location = "/apodtoday/?date=" + next;
}
document.onkeydown = function(e) {
    switch (e.keyCode) {
        case 37: // left
            document.getElementById("back").click();
            break;
        case 39: // right
            document.getElementById("forward").click();
            break;
    }
}
if (mediaType === "video") {
    videos();
}
function videos() {
    const iFrame =  document.getElementById("video-iframe");
    const imgCont = document.getElementById("img-cont");
    iFrame.style.display = "block";
    imgCont.style.display = "none";
    iFrame.src = imgCont.src;

}