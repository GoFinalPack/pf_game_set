//Global JS function for greeting
async function greet(link) {
    //Get user input
  const name = link.getAttribute('data');
  console.log(`Hello ${name}!`);
    //Call Go Greet function
    await window.go.main.App.SwitchPage(name,240,400);
}

// Project data
const projects = [
  { name: "Website Redesign", progress: 75, color: "bg-blue-500" },
  { name: "Mobile App Development", progress: 40, color: "bg-green-500" },
  { name: "Marketing Campaign", progress: 90, color: "bg-purple-500" },
  { name: "Database Migration", progress: 60, color: "bg-yellow-500" },
  { name: "AI Integration", progress: 25, color: "bg-red-500" },
  { name: "Customer Support Portal", progress: 80, color: "bg-indigo-500" }
];

// Function to create project cards
function createProjectCard(project) {
  return `
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
                    <h3 class="text-lg font-semibold text-gray-800 dark:text-white mb-2">${project.name}</h3>
                    <div class="flex items-center mb-4">
                        <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2.5">
                            <div class="h-2.5 rounded-full ${project.color}" style="width: ${project.progress}%"></div>
                        </div>
                        <span class="text-sm font-medium text-gray-500 dark:text-gray-400 ml-2">${project.progress}%</span>
                    </div>
                    <button class="text-blue-500 hover:text-blue-600 dark:text-blue-400 dark:hover:text-blue-300 text-sm font-medium">
                        View Details
                    </button>
                </div>
            `;
}

function createGameCard(game) {
  return ``;
}
// Toggle dark mode
function toggleDarkMode() {
  document.documentElement.classList.toggle('dark');
}
