<script>
    import { languages, currentLanguage } from '$lib/stores/language';
    import { onMount } from 'svelte';

    let currentLang = 'en';
    let isOpen = false;

    // Subscribe to the language store
    const unsubscribe = currentLanguage.subscribe(value => {
        currentLang = value;
    });

    /**
     * @param {MouseEvent} event
     * @param {string} targetId
     */
    function smoothScroll(event, targetId) {
        event.preventDefault();
        const element = document.getElementById(targetId);
        if (element) {
            element.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }
    }

    function toggleLanguage() {
        currentLanguage.update(lang => lang === 'en' ? 'bn' : 'en');
        isOpen = false;
    }

    // Close dropdown when clicking outside
    function handleClickOutside(event) {
        const target = event.target;
        if (!target.closest('.language-switcher')) {
            isOpen = false;
        }
    }

    onMount(() => {
        document.addEventListener('click', handleClickOutside);
        return () => {
            document.removeEventListener('click', handleClickOutside);
            unsubscribe();
        };
    });
</script>

<nav class="navbar">
    <div class="logo">
        <span class="logo-text">Dolil Lekhok</span>
    </div>
    <ul class="nav-links">
        <li><a href="#features" on:click={(e) => smoothScroll(e, 'features')}>Features</a></li>
        <li><a href="#how" on:click={(e) => smoothScroll(e, 'how')}>How It Works</a></li>
        <li><a href="#pricing" on:click={(e) => smoothScroll(e, 'pricing')}>Pricing</a></li>
    </ul>
    <div class="actions">
        <div class="language-switcher">
            <button 
                class="btn-secondary language-toggle"
                on:click|stopPropagation={() => isOpen = !isOpen}
                aria-haspopup="true"
                aria-expanded={isOpen}
                aria-label="Select language"
            >
                {currentLang === 'en' ? 'EN' : 'বাং'}
            </button>
            {#if isOpen}
                <div class="language-dropdown">
                    <button 
                        class="language-option {currentLang === 'en' ? 'active' : ''}"
                        on:click|stopPropagation={() => { toggleLanguage(); isOpen = false; }}
                    >
                        English
                    </button>
                    <button 
                        class="language-option {currentLang === 'bn' ? 'active' : ''}"
                        on:click|stopPropagation={() => { toggleLanguage(); isOpen = false; }}
                    >
                        বাংলা
                    </button>
                </div>
            {/if}
        </div>
        <a href="/login" class="btn-secondary">Login</a>
        <a href="/signup" class="btn-primary">Get Started</a>
    </div>
</nav>

<style>
.logo {
    display: flex;
    align-items: center;
    font-size: 1.5rem;
    font-weight: 700;
    letter-spacing: 0.5px;
}

.logo-text {
    background: linear-gradient(90deg, #8B5CF6 0%, #6366F1 100%);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    position: relative;
    transition: all 0.3s ease;
}

.logo-text::after {
    content: '';
    position: absolute;
    width: 100%;
    height: 2px;
    bottom: -4px;
    left: 0;
    background: linear-gradient(90deg, #8B5CF6 0%, #6366F1 100%);
    transform: scaleX(0);
    transform-origin: right;
    transition: transform 0.3s ease;
}

.logo:hover .logo-text::after {
    transform: scaleX(1);
    transform-origin: left;
}

.logo:hover .logo-text {
    text-shadow: 0 0 8px rgba(139, 92, 246, 0.3);
}

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
    align-items: center;
    position: relative;
}

.language-switcher {
    position: relative;
}

.language-toggle {
    min-width: 60px;
    text-align: center;
    padding: 0.5rem 1rem;
    cursor: pointer;
}

.language-dropdown {
    position: absolute;
    top: 100%;
    left: 0;
    background: #1E293B;
    border: 1px solid #334155;
    border-radius: 6px;
    margin-top: 0.5rem;
    min-width: 120px;
    z-index: 1000;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
    display: flex;
    flex-direction: column;
    overflow: hidden;
}

.language-option {
    padding: 0.5rem 1rem;
    background: transparent;
    border: none;
    color: #E2E8F0;
    text-align: left;
    cursor: pointer;
    transition: all 0.2s ease;
}

.language-option:hover {
    background-color: rgba(255, 255, 255, 0.05);
    color: #8B5CF6;
}

.language-option.active {
    color: #8B5CF6;
    font-weight: 500;
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
