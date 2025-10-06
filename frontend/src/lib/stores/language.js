import { writable } from 'svelte/store';

export const languages = {
    en: 'English',
    bn: 'বাংলা'
};

// Get initial language from localStorage or use default
const initialLang = (() => {
    if (typeof window === 'undefined') return 'en';
    const savedLang = localStorage.getItem('preferredLanguage');
    return savedLang && languages[savedLang] ? savedLang : 'en';
})();

// Create a custom store to handle language changes
export const currentLanguage = (() => {
    const { subscribe, set, update } = writable(initialLang);
    
    // Set up the HTML lang attribute
    if (typeof window !== 'undefined') {
        document.documentElement.lang = initialLang;
    }
    
    return {
        subscribe,
        set: (newLang) => {
            if (languages[newLang] && newLang !== initialLang) {
                localStorage.setItem('preferredLanguage', newLang);
                // Only reload if the language has actually changed
                if (typeof window !== 'undefined' && document.documentElement.lang !== newLang) {
                    document.documentElement.lang = newLang;
                    window.location.reload();
                }
            }
            set(newLang);
        },
        update: (updater) => {
            update(current => {
                const newLang = updater(current);
                if (languages[newLang] && newLang !== current) {
                    localStorage.setItem('preferredLanguage', newLang);
                    if (typeof window !== 'undefined' && document.documentElement.lang !== newLang) {
                        document.documentElement.lang = newLang;
                        window.location.reload();
                    }
                }
                return newLang;
            });
        }
    };
})();

export function toggleLanguage() {
    currentLanguage.update(lang => lang === 'en' ? 'bn' : 'en');
}

export function setLanguage(lang) {
    if (languages[lang]) {
        currentLanguage.set(lang);
    }
}
