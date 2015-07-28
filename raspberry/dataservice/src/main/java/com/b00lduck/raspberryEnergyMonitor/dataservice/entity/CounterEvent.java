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
public class CounterEvent {

	@Id
	@GeneratedValue(strategy = GenerationType.TABLE)
	private Long id;

	@ManyToOne
	@JoinColumn(name="counter_id", nullable = false)
	private Counter counter;

	private DateTime timestamp;

	private CounterEventType type;

	private BigDecimal delta;

	private BigDecimal reading;

	public Long getId() {
		return id;
	}

	public Counter getCounter() {
		return counter;
	}

	public void setCounter(final Counter counter) {
		this.counter = counter;
	}

	public CounterEventType getType() {
		return type;
	}

	public void setType(final CounterEventType type) {
		this.type = type;
	}

	public DateTime getTimestamp() {
		return timestamp;
	}

	public void setTimestamp(final DateTime timestamp) {
		this.timestamp = timestamp;
	}

	public BigDecimal getDelta() {
		return delta;
	}

	public void setDelta(BigDecimal delta) {
		this.delta = delta;
	}

	public BigDecimal getReading() {
		return reading;
	}

	public void setReading(BigDecimal reading) {
		this.reading = reading;
	}

}
