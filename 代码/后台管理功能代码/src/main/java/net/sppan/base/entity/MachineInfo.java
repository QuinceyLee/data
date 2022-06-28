package net.sppan.base.entity;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import net.sppan.base.entity.support.BaseEntity;

import javax.persistence.*;
import java.util.Date;


@Entity
@Table(name = "machine_info")
@JsonIgnoreProperties({"hibernateLazyInitializer","handler"})
@Data
public class MachineInfo extends BaseEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", nullable = false)
    private Integer id;

    /**
     * 角色名称
     */
    private String number;

    /**
     * 角色key
     */
    private String name;

    /**
     * 角色状态,0：正常；1：删除
     */
    private String type;

    /**
     * 角色描述
     */
    private String location;


    private String supplier;

    private String personId;

    private String UUID;

}
