IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.tKegSurvei')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.tKegSurvei (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            survei_kd varchar(50) NOT NULL, -- references?
            keg_kd    varchar(10) NOT NULL, -- references?
            /* reference to another table example
            UserID  int FOREIGN KEY REFERENCES Users(ID),*/
            status    tinyint NOT NULL,
            tgl_buka  datetime,
            tgl_rek_mulai datetime,
            tgl_rek_selesai datetime,
            tgl_mulai datetime,
            tgl_selesai datetime,
            is_rekrutmen tinyint,
            is_multi  tinyint,
            is_confirm bit NOT NULL,
            is_add_indicator bit NOT NULL,
            /*versi_will tinyint,
            is_cocard tinyint,
            is_cocard_published tinyint,
            jum_struktur tinyint,
            tipe_penugasan tinyint,
            max_posisi_sampel tinyint,
            is_assign_struktur tinyint,
            is_assign_sampel tinyint,
            min_delete_penugasan smallint,*/
            --create+update
            created_by varchar(20) NOT NULL,
            created_at datetime default CURRENT_TIMESTAMP,
            updated_by varchar(20),
            updated_at datetime,
        );
/*GO

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
        );
        
        
        IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.tPetugasSurveiBranch')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.tPetugasSurveiBranch (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            survei_kd varchar(20) NOT NULL, -- references?
            keg_kd    varchar(10) NOT NULL, -- references?
            branch_kd varchar(20) NOT NULL, -- references?
            posisi_kd varchar(10) NOT NULL, -- references
            username  varchar(20) FOREIGN KEY REFERENCES tUser(username),
            status tinyint,
            --create+update
        );


        IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.tPetugasKinerjaSurvei')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.tPetugasKinerjaSurvei (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            survei_kd varchar(20) NOT NULL, -- references?
            keg_kd    varchar(10) NOT NULL, -- references?
            branch_kd varchar(20) NOT NULL, -- references?
            posisi_kd varchar(10) NOT NULL, -- references
            penilai   varchar(20) NOT NULL,
            nilai     tinyint,
            --create+update
        );
        
        IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.penugasanStruktur')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.penugasanStruktur (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            keg_kd    varchar(10) NOT NULL, -- references?
            branch_kd varchar(20) NOT NULL, -- references?
            username  varchar(20) FOREIGN KEY REFERENCES tUser(username),
            posisi_kd varchar(10) NOT NULL, -- references
            parent   varchar(20) NOT NULL,
            status     tinyint,
            --create+update
        );
        
        
        
        
        IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.nilai')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.penugasanStruktur (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            survei_kd varchar(20) NOT NULL, -- references?
            keg_kd    varchar(10) NOT NULL, -- references?
            branch_kd varchar(20) NOT NULL, -- references?
            indikator_id  varchar(20) FOREIGN KEY REFERENCES indikator(ID),
            nilai     tinyint,
            --create+update
        );


        IF NOT EXISTS
        (SELECT *FROM sysobjects
            WHERE id = object_id(N'dbo.indikator')
            AND OBJECTPROPERTY(id, N'IsUserTable') = 1
        )
        CREATE TABLE dbo.indikator (
            ID        int IDENTITY(1,1) PRIMARY KEY,
            branch_kd varchar(20) NOT NULL, -- references?
            indikator varchar(20) NOT NULL,
            definisi  varchar(255) NOT NULL,
            bobot     tinyint NOT NULL,
            label     varchar(255) NOT NULL,
            --create+update
        );*/