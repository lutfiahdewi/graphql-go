IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.tPosKegSurvei')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.tPosKegSurvei (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            survei_kd varchar(20) NOT NULL, -- references?
            keg_kd    varchar(20) NOT NULL, -- references?
            /* reference to another table example
            UserID  int FOREIGN KEY REFERENCES Users(ID),*/
           urutan     tinyint,
           is_view_mitra tinyint,
           is_multi tinyint,
           --create+update
            created_by varchar(20) NOT NULL,
            created_at datetime default CURRENT_TIMESTAMP,
            updated_by varchar(20),
            updated_at datetime,
        );
        