package net.sppan.base.entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import net.sppan.base.entity.support.BaseEntity;

import javax.persistence.*;
@Entity
@Table(name = "gateway_info")
@JsonIgnoreProperties({"hibernateLazyInitializer","handler"})
@Data
public class GatewayInfo extends BaseEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", nullable = false)
    private Integer id;

    /**
     * 角色名称
     */
    private String uuid;

    /**
     * 角色key
     */
    private String status;

    /**
     * 角色状态,0：正常；1：删除
     */
    private String description;

}



