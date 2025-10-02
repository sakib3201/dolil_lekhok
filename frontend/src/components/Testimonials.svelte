<script>
    import { fade } from 'svelte/transition';
    import demo_profile_1 from '../images/demo_profile_1.webp';
    import demo_profile_2 from '../images/demo_profile_2.webp';
    import demo_profile_3 from '../images/demo_profile_3.webp';
</script>

<section class="testimonials-section" in:fade>
    <h2 class="section-title">বাংলাদেশী আইনজীবীদের নিকট বিশ্বাসযোগ্য</h2>
    <p class="section-subtitle">Dolil Lekhok সম্পর্কে আইনজীবীদের মতামত দেখুন</p>
    <div class="testimonials-grid">
        <div class="testimonial-card">
            <img src={demo_profile_1} alt="Advocate Rahman" class="testimonial-image">
            <h3>অ্যাডভোকেট মানিলা</h3>
            <p class="testimonial-text">"এই সফটওয়্যারটি আমাদের আইনি কাজে বিপ্লব ঘটিয়েছে। নির্ভুলতা অসাধারণ!"</p>
            <div class="rating">★★★★★</div>
        </div>
        <div class="testimonial-card">
            <img src={demo_profile_2} alt="Barrister Khan" class="testimonial-image">
            <h3>ব্যারিস্টার খান</h3>
            <p class="testimonial-text">"আইনি দলিল ফরম্যাটিংয়ের জন্য এটি পারফেক্ট। আমি সুপারিশ করি।"</p>
            <div class="rating">★★★★☆</div>
        </div>
        <div class="testimonial-card">
            <img src={demo_profile_3} alt="Legal Clerk Ahmed" class="testimonial-image">
            <h3>ব্যারিস্টার আহমেদ</h3>
            <p class="testimonial-text">"একটি সহজে ব্যবহারযোগ্য টুল যা আমার কাজে বিরাট সুবিধা দেয়।"</p>
            <div class="rating">★★★★★</div>
        </div>
    </div>
</section>

<style>
    .testimonials-section {
        padding: 80px 20px;
        background-color: #0F172A;
        color: #FFFFFF;
        text-align: center;
        position: relative;
        overflow: hidden;
    }
    
    .testimonials-section::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 1px;
        background: linear-gradient(90deg, transparent, rgba(99, 102, 241, 0.5), transparent);
    }
    .section-title {
        font-size: 2.5rem;
        font-weight: 600;
        margin-bottom: 1rem;
        color: #6366F1;
        text-shadow: 0 2px 4px rgba(0,0,0,0.1);
        position: relative;
        display: inline-block;
    }
    
    .section-title::after {
        content: '';
        position: absolute;
        bottom: -10px;
        left: 50%;
        transform: translateX(-50%);
        width: 60px;
        height: 3px;
        background: #8B5CF6;
        border-radius: 2px;
    }
    .section-subtitle {
        font-size: 1.125rem;
        font-weight: 400;
        color: #94A3B8;
        margin-bottom: 4rem;
        max-width: 700px;
        margin-left: auto;
        margin-right: auto;
        line-height: 1.6;
    }
    .testimonials-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
        gap: 2rem;
        padding: 0 20px;
        max-width: 1200px;
        margin: 0 auto;
        position: relative;
    }
    .testimonial-card {
        background: linear-gradient(145deg, #1E293B, #0F172A);
        padding: 2rem;
        border: 1px solid #334155;
        border-radius: 12px;
        transition: all 0.3s ease;
        position: relative;
        overflow: hidden;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    }
    
    .testimonial-card::before {
        content: '"';
        position: absolute;
        top: 1rem;
        right: 1.5rem;
        font-size: 4rem;
        color: rgba(99, 102, 241, 0.1);
        font-family: serif;
        line-height: 1;
    }
    .testimonial-card:hover {
        transform: translateY(-8px);
        border-color: #6366F1;
        box-shadow: 0 10px 25px -5px rgba(99, 102, 241, 0.1);
    }
    .testimonial-image {
        width: 80px;
        height: 80px;
        border-radius: 50%;
        margin-bottom: 1.5rem;
        object-fit: cover;
        border: 3px solid #6366F1;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
        transition: all 0.3s ease;
    }
    
    .testimonial-card:hover .testimonial-image {
        transform: scale(1.1);
        border-color: #8B5CF6;
    }
    .testimonial-card h3 {
        font-size: 1.25rem;
        font-weight: 600;
        margin-bottom: 0.5rem;
        color: #FFFFFF;
        position: relative;
        display: inline-block;
    }
    
    .testimonial-card h3::after {
        content: '';
        position: absolute;
        bottom: -0.5rem;
        left: 0;
        width: 30px;
        height: 2px;
        background: #8B5CF6;
        transition: width 0.3s ease;
    }
    
    .testimonial-card:hover h3::after {
        width: 50px;
    }
    .testimonial-text {
        font-size: 1rem;
        line-height: 1.7;
        font-weight: 400;
        color: #94A3B8;
        margin: 1.5rem 0;
        position: relative;
        z-index: 1;
    }
    .rating {
        font-size: 1.25rem;
        color: #F59E0B;
        letter-spacing: 2px;
        margin-top: 1rem;
    }
</style>