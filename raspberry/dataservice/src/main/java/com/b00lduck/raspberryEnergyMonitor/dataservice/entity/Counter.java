package com.b00lduck.raspberryEnergyMonitor.dataservice.entity;

import org.hibernate.annotations.Type;
import org.joda.time.DateTime;

import javax.persistence.*;
import java.math.BigDecimal;

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

	private BigDecimal reading;

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

	public DateTime getLastTick() {
		return lastTick;
	}

	public void setLastTick(final DateTime lastTick) {
		this.lastTick = lastTick;
	}

	public BigDecimal getReading() {
		return reading;
	}

	public void setReading(BigDecimal reading) {
		this.reading = reading;
	}

}
