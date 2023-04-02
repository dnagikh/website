document.getElementById("fooDropdownBtn").addEventListener("click", function (event) {
    document.getElementById("fooDropdown").classList.toggle("hidden");
    event.target.getElementsByTagName('svg')[0].classList.toggle("rotate-180");
})