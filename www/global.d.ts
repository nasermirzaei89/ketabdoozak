import htmx from "htmx.org";
import { Alpine as AlpineType } from 'alpinejs'

export {}; // Ensure this file is treated as a module

declare global {
    interface Window {
        chooseThumbnailUrl: () => void;
        htmx: typeof htmx;
        Alpine: AlpineType;
    }
}
