import { format, formatDistanceToNow } from 'date-fns';

export function ParseDate(dateString: string): string {
    const date = new Date(dateString);

    // For "1 hour ago"
    const relative = formatDistanceToNow(date, { addSuffix: true });
    
    // For "14 Sep 2025 15:00PM"
    const formatted = format(date, 'dd MMM yyyy HH:mm') + 'PM';
    return relative
}
