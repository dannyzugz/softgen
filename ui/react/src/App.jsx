import "./App.css";

function App() {

    return (
        <main>
            <h1>Code Generator</h1>
            <h2>STEP 1 Choose your application type</h2>

            <form action="">
                <div>
                    <label>
                        <input type="radio" name="app-type" />
                        Traditional web application Generate the scaffold for a
                        'traditional' web application which uses server-side
                        HTML rendering.
                    </label>
                </div>
                <div>
                    <label>
                        <input type="radio" name="app-type"/>
                        JSON API Generate the scaffold for a API which processes
                        JSON requests and sends JSON responses.
                    </label>
                </div>
            </form>

            <h2> STEP 2 Customize your application</h2>

            <strong>Choose a database:</strong>

            <form action="">
                <div>
                    <label>
                        <input type="radio" name="database-type"/>
                        None
                    </label>
                    <label>
                        <input type="radio" name="database-type"/>
                        PostgreSQL
                    </label>
                    <label>
                        <input type="radio" name="database-type"/>
                        MySQL
                    </label>
                </div>
                
                <strong>Pick your preferred router:</strong>

                <div>
                    <label>
                        <input type="radio" name="router-type"/>
                        Chi
                    </label>
                    <label>
                        <input type="radio" name="router-type"/>
                        Flow
                    </label>
                    <label>
                        <input type="radio" name="router-type"/>
                        Gorila Mux
                    </label>
                    <label>
                        <input type="radio" name="router-type"/>
                        HttpRouter
                    </label>
                    <label>
                        <input type="radio" name="router-type"/>
                        Gin 
                    </label>
                </div>
            </form>
        </main>
    );
}

export default App;


// Command-line flags
// Environment variables
// Choose your log format:

// Plaintext
// JSON
// Colorized plaintext
// Enter the module path for your project:

// your.module/path
// STEP 3

// Select additional features and functionality
// Tick any you want to include:

// Access logging – Middleware for logging all requests and corresponding response status code and body size.

// Admin Makefile – Makefile actions for building binaries and vetting, testing and linting your code.

// Automatic versioning – Use your VCS revision identifier (e.g. Git commit hash) as the application version number.

// Basic auth – Middleware for adding HTTP basic authentication to your application endpoints.

// Emails – Helpers for building and sending emails via SMTP.

// Error notifications – Send runtime error alerts to an admin email address.

// Gitignore – Initialize a .gitignore file for the project.

// Live reload – Watch your files for changes and automatically rebuild and restart your application during development.

// Secure cookies – Helpers for reading/writing signed and encrypted cookies.

// Sessions – Client-side sessions using the gorilla/sessions package.

// SQL migrations – Makefile actions for managing SQL migrations, and automatically run migrations on application startup.
// To use PLUS features you'll need an Autostrada Plus account, available for a one-time purchase of $29. All other features are free!

// Automatic HTTPS – Serve requests over HTTPS when running in production. Automatic TLS certificate management via Let's Encrypt and HTTP to HTTPS redirects. PLUS

// Custom error pages – Use custom pages for error responses (e.g. 404 Not Found and 500 Internal Server Error pages). PLUS

// User accounts and authentication – User accounts with email/password authentication and fully-functional signup, login, logout and password reset workflows. Authorization middleware and CSRF protection also included. PLUS
// CUSTOMIZATION COMPLETE!

// Click the button below to generate a ZIP file containing your project code.

// Generate your project code!
