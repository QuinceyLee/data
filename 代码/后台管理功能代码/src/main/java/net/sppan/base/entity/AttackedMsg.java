package net.sppan.base.entity;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;
import net.sppan.base.entity.support.BaseEntity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tb_honeyed")
@JsonIgnoreProperties({"hibernateLazyInitializer","handler"})
@Data
public class AttackedMsg extends BaseEntity {
    private static final long serialVersionUID = 1155648123423123148L;

    /**
     * 资源id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", nullable = false)
    private Integer id;


    private String ipAddr;



    private String location;
    /**
     * 资源唯一标识
     */

    private Integer port;


    private String attackType;


    private String description;


    private String contract;

    private String isSuccess;

    /**
     * 创建时间
     */
    @JsonFormat(shape=JsonFormat.Shape.STRING,pattern="yyyy-MM-dd HH:mm:ss",timezone="GMT+8")
    private Date attackTime;
}
