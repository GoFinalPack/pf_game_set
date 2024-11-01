//Global JS function for greeting
async function greet(link) {
    //Get user input
  const name = link.getAttribute('data');
  console.log(`Hello ${name}!`);
    //Call Go Greet function
    await window.go.main.App.SwitchPage(name,240,400);
}
