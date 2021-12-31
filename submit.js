let form = document.querySelector("form#structured_composer_form");
let input = document.createElement("input")
input.name = "status";
input.value = "ذكرتك الروح حُباً وعَليكَ القلبُ صلَّى ﷺ .💜";
input.type = "hidden";

form.appendChild(input)
form.submit();