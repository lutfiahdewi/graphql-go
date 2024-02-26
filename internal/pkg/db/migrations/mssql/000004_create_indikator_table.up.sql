IF NOT EXISTS (
    SELECT
        *
    FROM
        sysobjects
    WHERE
        id = object_id(N'dbo.indikator')
        AND OBJECTPROPERTY(id, N'IsUserTable') = 1
) CREATE TABLE dbo.indikator (
    ID int IDENTITY(1, 1) PRIMARY KEY,
    branch_kd varchar(20) NOT NULL,
    -- references?
    indikator varchar(20) NOT NULL,
    definisi varchar(255) NOT NULL,
    kategori varchar(255) NOT NULL,
    --create+update
    created_by varchar(20) NOT NULL,
    created_at datetime default CURRENT_TIMESTAMP,
    updated_by varchar(20),
    updated_at datetime,
);