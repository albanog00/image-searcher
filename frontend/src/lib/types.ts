export enum SearchState {
    Default,
    Finished,
    Searching,
    Error,
}

export type SearchPaginationQuery = {
    page: number;
    perPage: number;
    autoPaging: boolean;
}

export interface Result<T> {
    message: string,
    status: number,
    result?: {
        data: T,
        count: number
    } | null
}

export type ProfileImage = {
    small: string;
    medium: string;
    large: string;
}

export type UserLinks = {
    self: string;
    html: string;
    photos: string;
    likes: string;
}

export type User = {
    id: string;
    username: string;
    name: string;
    first_name: string;
    last_name: string;
    instagram_username: string;
    portfolio_url: string;
    profile_image: ProfileImage;
    links: UserLinks;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type Collection = any;

export type Urls = {
    raw: string;
    full: string;
    regular: string;
    small: string;
    thumb: string;
}

export type PhotoLinks = {
    self: string;
    html: string;
    download: string;
}

export type Photo = {
    id: string;
    created_at: Date;
    updated_at: Date;
    widht: number;
    height: number;
    color: string;
    blur_hash: string;
    likes: number;
    liked_by_user: boolean;
    description: string;
    user: User;
    current_user_collections: Collection[];
    urls: Urls;
    links: PhotoLinks
}

export type SearchPhotoResult = {
    total: number;
    total_pages: number;
    results: Photo[];
}
