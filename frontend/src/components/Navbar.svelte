<nav class="navbar">
    <div class="logo">Dolil Lekhok</div>
    <ul class="nav-links">
        <li><a href="#features">Features</a></li>
        <li><a href="#how">How It Works</a></li>
        <li><a href="#pricing">Pricing</a></li>
        <li><a href="#contact">Contact</a></li>
    </ul>
    <div class="actions">
        <a href="/login" class="btn-secondary">Login</a>
        <a href="/signup" class="btn-primary">Get Started</a>
    </div>
</nav>

<style>
.navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    background: #1E293B; /* Dark theme background */
    color: #FFFFFF;
    border-bottom: 1px solid #334155;
}

.nav-links {
    list-style: none;
    display: flex;
    gap: 2rem;
}

.nav-links a {
    color: #E2E8F0;
    text-decoration: none;
    font-weight: 500;
    transition: color 0.2s ease;
}

.nav-links a:hover {
    color: #8B5CF6; /* Secondary color */
}

.actions {
    display: flex;
    gap: 1rem;
}

.btn-primary {
    background-color: #6366F1; /* Primary color */
    color: white;
    padding: 0.5rem 1.5rem;
    border-radius: 6px;
    font-weight: 500;
    transition: background-color 0.2s ease;
}

.btn-primary:hover {
    background-color: #4F46E5;
}

.btn-secondary {
    background-color: transparent;
    color: #E2E8F0;
    padding: 0.5rem 1.5rem;
    border: 1px solid #4B5563;
    border-radius: 6px;
    font-weight: 500;
    transition: all 0.2s ease;
}

.btn-secondary:hover {
    background-color: rgba(255, 255, 255, 0.05);
    border-color: #8B5CF6; /* Secondary color */
    color: #8B5CF6; /* Secondary color */
}
</style>
