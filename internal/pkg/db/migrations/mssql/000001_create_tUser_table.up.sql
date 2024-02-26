IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.tUser')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.tUser (
            ID         int IDENTITY(1,1) PRIMARY KEY,
            username   varchar(20) NOT NULL,
            password   varchar(255) NOT NULL,
            nik        varchar(20) NOT NULL,
            nip        varchar(20) NOT NULL,
            sobat_id   varchar(20) NOT NULL,
            email      varchar(20) NOT NULL,
            nama       varchar(255) NOT NULL,
            status     tinyint,
            status_blacklist tinyint,
            no_hp      varchar(20),
            is_agree   tinyint,
            is_pegawai tinyint,
            unit_id    int,
            status_nik tinyint,
            created_by varchar(20) NOT NULL,
            created_at datetime default CURRENT_TIMESTAMP,
            updated_by varchar(20),
            updated_at datetime,
        );