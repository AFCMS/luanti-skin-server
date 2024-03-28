declare namespace ApiTypes {
    type UserResponse = {
        id: number;
        name: string;
        email: string;
        permission_level: 1 | 2 | 3 | 4;
        banned?: boolean;
        ban_reason?: string;
        created_at: string;
        last_connection: string;
    };

    type SkinResponse = {
        uuid: string;
        description: string;
        public?: boolean;
        approved?: boolean;
        owner_id: number;
        creation_date: string;
    };

    type AccountUserResponse = {
        id: number;
        username: string;
        permission_level: 1 | 2 | 3 | 4;
        cdb_username: string;
    };

    type InfoProviderTypes = "contentdb" | "github" | "codeberg" | "discord";

    type InfoResponse = {
        account_count: number;
        skin_count: number;
        version: string;
        supported_oauth_providers: InfoProviderTypes[];
    };
}
