CREATE DATABASE IF NOT EXISTS gestion_usuarios;
USE gestion_usuarios;
-- Tabla de Usuarios
CREATE TABLE usuarios (
                          usuario_id CHAR(36) PRIMARY KEY,
                          nombre VARCHAR(100) NOT NULL,
                          apellido VARCHAR(100) NOT NULL,
                          email VARCHAR(255) NOT NULL UNIQUE,
                          contrasena VARCHAR(255) NOT NULL,
                          fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          activo BOOLEAN DEFAULT TRUE
);

-- Tabla de Organizaciones
CREATE TABLE organizaciones (
                                organizacion_id CHAR(36) PRIMARY KEY,
                                nombre VARCHAR(255) NOT NULL UNIQUE,
                                descripcion TEXT NULL,
                                fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                activo BOOLEAN DEFAULT TRUE
);

-- Tabla de Organización por Usuario (Relación N:M)
CREATE TABLE organizacion_por_usuario (
                                          organizacion_por_usuario_id CHAR(36) PRIMARY KEY,
                                          organizacion_id CHAR(36) NOT NULL,
                                          usuario_id CHAR(36) NOT NULL,
                                          FOREIGN KEY (organizacion_id) REFERENCES organizaciones(organizacion_id),
                                          FOREIGN KEY (usuario_id) REFERENCES usuarios(usuario_id)
);

-- Tabla de Roles
CREATE TABLE roles (
                       rol_id CHAR(36) PRIMARY KEY,
                       nombre VARCHAR(100) NOT NULL UNIQUE,
                       descripcion TEXT NULL,
                       activo BOOLEAN DEFAULT TRUE
);

-- Tabla de Usuarios por Roles (Relación N:M)
CREATE TABLE usuarios_por_roles (
                                    usuario_por_rol_id CHAR(36) PRIMARY KEY,
                                    usuario_id CHAR(36) NOT NULL,
                                    rol_id CHAR(36) NOT NULL,
                                    FOREIGN KEY (usuario_id) REFERENCES usuarios(usuario_id),
                                    FOREIGN KEY (rol_id) REFERENCES roles(rol_id)
);

-- Tabla de Permisos
CREATE TABLE permisos (
                          permiso_id CHAR(36) PRIMARY KEY,
                          nombre VARCHAR(100) NOT NULL UNIQUE,
                          descripcion TEXT NULL,
                          activo BOOLEAN DEFAULT TRUE
);

-- Tabla de Roles por Permisos (Relación N:M)
CREATE TABLE roles_por_permisos (
                                    rol_por_permiso_id CHAR(36) PRIMARY KEY,
                                    rol_id CHAR(36) NOT NULL,
                                    permiso_id CHAR(36) NOT NULL,
                                    FOREIGN KEY (rol_id) REFERENCES roles(rol_id),
                                    FOREIGN KEY (permiso_id) REFERENCES permisos(permiso_id)
);
