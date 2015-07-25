package com.b00lduck.raspberryEnergyMonitor.dataservice.entity;

import org.hibernate.annotations.Type;
import org.hibernate.annotations.UpdateTimestamp;
import org.joda.time.DateTime;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import javax.persistence.*;
import java.math.BigDecimal;
import java.util.List;

/**
 * @author Daniel Zerlett (daniel@zerlett.eu)
 * created 24.07.2015
 */
@Entity
public class Counter {

	@Id
	@GeneratedValue(strategy = GenerationType.TABLE)
	private Long id;

	private String name;

	private String unit;

	private BigDecimal value;

	@Type(type="org.jadira.usertype.dateandtime.joda.PersistentDateTime")
	private DateTime lastTick;

    public Long getId() {
		return id;
	}

	public String getName() {
		return name;
	}

	public void setName(final String name) {
		this.name = name;
	}

	public String getUnit() {
		return unit;
	}

	public void setUnit(final String unit) {
		this.unit = unit;
	}

	public BigDecimal getValue() {
		return value;
	}

	public void setValue(final BigDecimal value) {
		this.value = value;
	}

	public DateTime getLastTick() {
		return lastTick;
	}

	public void setLastTick(final DateTime lastTick) {
		this.lastTick = lastTick;
	}

}
