<script>
    import { layout_images } from "@stores/layout";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import HeroVideo from "../sub-components/hero_video.svelte";
</script>

<section id="smk-hero-section" class:adebug={false}>
    <header id="smk-hp-header">
        <div id="combination-mark-wrapper">
            <img src="{layout_images.COMBINATION_MARK.getUrl(0.13)}" alt="selvamaaku combination mark">
        </div>
        <div id="smk-hero-headline-wrapper">
            <div class="wixarica-decoration">
                <WixaricaIcon opacity={0.4}/>
            </div>
            <h1 id="smk-hero-headline">
                Living where Jungle Bliss<br/>Meets Coastal Luxury
            </h1>
            <div class="wixarica-decoration">
                <WixaricaIcon opacity={0.4}/>
            </div>
        </div>
        <div id="smk-hero-subheadline-wrapper">
            <p id="smk-hero-subheadline">
                Live in the most exclusive house development in the jungle outside Sayulita, Mexico.
            </p>
        </div>
    </header>
    <div id="hero-video-wrapper" class:adebug={false}>
        <HeroVideo
            video_width="71.11cqw"
            video_height="40cqw"
        />
    </div>
</section>

<style>
    header#smk-hp-header {
        display: flex;
        flex-direction: column;
        padding: var(--spacing-4) 0 0 0;
        row-gap: var(--spacing-3);
        align-items: center;
    }

    
    /*=============================================
    =            Combination mark            =
    =============================================*/
    
        #combination-mark-wrapper {
            width: 100%;
            display: grid;
            place-items: center;
        }

        #combination-mark-wrapper img {
            width: 242px;
        }
    
    /*=====  End of Combination mark  ======*/

    
    /*=============================================
    =            Headline            =
    =============================================*/
    
        #smk-hero-headline-wrapper {
            display: flex;
            width: max-content;
            gap: var(--spacing-4);
            align-items: center;
            justify-content: center;
        }

        .wixarica-decoration {
            width: 120px;
        }

        h1#smk-hero-headline {
            font-size: var(--font-size-h2);
            text-align: center;
        }

        #smk-hero-subheadline-wrapper {
            width: max-content;
        } 

        p#smk-hero-subheadline {
            max-width: 527px;
            font-family: var(--font-read);
            font-size: var(--font-size-2);
            text-align: center;
        }
    
    /*=====  End of Headline  ======*/
    
    
    /*=============================================
    =            Video Section            =
    =============================================*/
    
        #hero-video-wrapper {
            display: flex;
            container-type: inline-size;
            width: min(100%, var(--design-width));
            justify-content: center;
            padding: var(--spacing-4) 0;
            margin: 0 auto;
        }
    
    /*=====  End of Video Section  ======*/
    
    
    
    
</style>