<script>
    import { goto } from '$app/navigation';
    
    let email = '';
    let password = '';
    let confirmPassword = '';
    let error = '';
    
    async function handleSignup(e) {
        e.preventDefault();
        error = '';
        
        if (password !== confirmPassword) {
            error = 'Passwords do not match';
            return;
        }
        
        try {
            const response = await fetch('http://localhost:8080/api/auth/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: email,
                    password: password
                })
            });
            
            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.error || 'Signup failed');
            }
            
            // Redirect to login page after successful signup
            goto('/login');
            
        } catch (err) {
            error = err.message || 'Signup failed. Please try again.';
        }
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 p-8 bg-gray-800 rounded-xl shadow-2xl">
        <div>
            <h2 class="mt-6 text-center text-3xl font-extrabold text-white">
                Create Your Account
            </h2>
            <p class="mt-2 text-center text-sm text-gray-400">
                Already have an account? 
                <a href="/login" class="font-medium text-indigo-400 hover:text-indigo-300 transition-colors">
                    Sign in
                </a>
            </p>
        </div>
        
        {#if error}
            <div class="bg-red-900/30 border border-red-700 text-red-200 px-4 py-3 rounded-lg relative" role="alert">
                <span class="block sm:inline">{error}</span>
            </div>
        {/if}
        
        <form class="mt-8 space-y-6" on:submit|preventDefault={handleSignup}>
            <div class="space-y-4">
                <div>
                    <label for="email-address" class="sr-only">Email address</label>
                    <div class="relative">
                        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                            <svg class="h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                                <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                                <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                            </svg>
                        </div>
                        <input
                            id="email-address"
                            name="email"
                            type="email"
                            autocomplete="email"
                            required
                            bind:value={email}
                            class="appearance-none block w-full pl-10 pr-3 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent sm:text-sm"
                            placeholder="Email address"
                        />
                    </div>
                </div>
                <div>
                    <label for="password" class="sr-only">Password</label>
                    <div class="relative">
                        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                            <svg class="h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                                <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                            </svg>
                        </div>
                        <input
                            id="password"
                            name="password"
                            type="password"
                            autocomplete="new-password"
                            required
                            bind:value={password}
                            class="appearance-none block w-full pl-10 pr-3 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent sm:text-sm"
                            placeholder="Password"
                        />
                    </div>
                </div>
                <div>
                    <label for="confirm-password" class="sr-only">Confirm Password</label>
                    <div class="relative">
                        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                            <svg class="h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                                <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                            </svg>
                        </div>
                        <input
                            name="confirm-password"
                            type="password"
                            autocomplete="new-password"
                            required
                        bind:value={confirmPassword}
                        class="appearance-none block w-full pl-10 pr-3 py-3 bg-gray-700 border border-gray-600 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent sm:text-sm"
                        placeholder="Confirm Password"
                    />
                </div>
            </div>

            <div class="flex items-start">
                <div class="flex items-center h-5">
                    <input
                        id="terms"
                        name="terms"
                        type="checkbox"
                        required
                        class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-600 rounded bg-gray-700"
                    />
                </div>
                <div class="ml-3 text-sm">
                    <label for="terms" class="text-gray-300">
                        I agree to the <a href="#" class="font-medium text-indigo-400 hover:text-indigo-300 transition-colors">Terms and Conditions</a>
                    </label>
                </div>
            </div>

            <div>
                <button
                    type="submit"
                    class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-lg text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors duration-200"
                >
                    <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                        <svg class="h-5 w-5 text-indigo-300 group-hover:text-indigo-200 transition-colors" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                            <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
                        </svg>
                    </span>
                    Create Account
                </button>
            </div>
        </form>
    </div>
</div>

<style>
    /* Smooth transitions for interactive elements */
    input, button, a {
        transition: all 0.2s ease-in-out;
    }
    
    /* Focus styles */
    input:focus, button:focus {
        outline: none;
        box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.3);
    }
    
    /* Animation for the form */
    @keyframes fadeIn {
        from { opacity: 0; transform: translateY(10px); }
        to { opacity: 1; transform: translateY(0); }
    }
    
    form {
        animation: fadeIn 0.5s ease-out;
    }
    
    /* Password strength indicator */
    .password-strength {
        height: 4px;
        margin-top: 0.5rem;
        border-radius: 2px;
        transition: all 0.3s ease;
    }
    
    .strength-weak { background-color: #ef4444; width: 33%; }
    .strength-medium { background-color: #f59e0b; width: 66%; }
    .strength-strong { background-color: #10b981; width: 100%; }
</style>
