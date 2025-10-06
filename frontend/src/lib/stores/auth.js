import { writable } from 'svelte/store';

export const user = writable(null);

export async function login(username, password) {
    const response = await fetch('http://localhost:8080/api/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
        credentials: 'include'
    });

    if (response.ok) {
        const data = await response.json();
        user.set({ username });
        return { success: true };
    }
    
    const error = await response.json();
    return { success: false, error: error.error || 'Login failed' };
}

export async function signup(username, password) {
    const response = await fetch('http://localhost:8080/api/auth/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password })
    });

    if (response.ok) {
        return { success: true };
    }
    
    const error = await response.json();
    return { success: false, error: error.error || 'Signup failed' };
}

export async function logout() {
    // If you implement server-side logout, call it here
    user.set(null);
    // Redirect to home page
    window.location.href = '/';
}

export async function checkAuth() {
    try {
        const response = await fetch('http://localhost:8080/api/auth/me', {
            credentials: 'include'
        });
        
        if (response.ok) {
            const data = await response.json();
            user.set(data.user);
            return true;
        }
        return false;
    } catch (error) {
        console.error('Auth check failed:', error);
        return false;
    }
}
